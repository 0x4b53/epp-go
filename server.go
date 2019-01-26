package eppserver

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/xml"
	"log"
	"net"
	"sync"
	"time"

	proxyproto "github.com/armon/go-proxyproto"
	"github.com/pkg/errors"
)

const (
	listenTimeout      = 1 * time.Second
	requestReadTimeout = 100 * time.Millisecond

	// RFC 5734, section 8
	contentLengthHeaderBytes = 4
)

// Server represents the server handling requests.
type Server struct {
	// VerifyClientCertificateFunc is the function that will be set in the TLS
	// config to verify the client certificate when connection to the EPP
	// server.
	VerifyClientCertificateFunc func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

	// Server ID will be used in the greeting message to say which server sent
	// the greeting. This should probably be the hostname of the EPP server.
	ServerID string

	keyFile  string
	certFile string
	listen   string
	stopChan chan struct{}
	wg       sync.WaitGroup
	mu       sync.RWMutex
}

// New returns a new instance of a EPP server.
func New(serverID, listen, keyFile, certFile string) *Server {
	return &Server{
		VerifyClientCertificateFunc: func(_ [][]byte, _ [][]*x509.Certificate) error {
			log.Print("YOU SHOULD IMPLEMENT VERIFICATION OF THE CLIENT CERTIFICATE")

			return errors.New("client certificate verification not implemented")
		},
		ServerID: serverID,
		listen:   listen,
		keyFile:  keyFile,
		certFile: certFile,
		stopChan: make(chan struct{}),
		wg:       sync.WaitGroup{},
		mu:       sync.RWMutex{},
	}
}

// ListenAndServe will start a new server and handle incomming connections.
func (s *Server) ListenAndServe() error {
	addr, err := net.ResolveTCPAddr("tcp", s.listen)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("listening for connections on %s", s.listen)

	defer func() {
		if err := listener.Close(); err != nil {
			log.Print(err.Error())
		}

		log.Print("waiting for all ongoing requests")
		s.wg.Wait()
	}()

	cert, err := tls.LoadX509KeyPair(s.certFile, s.keyFile)
	if err != nil {
		log.Fatal("Error loading certificate. ", err)
	}

	tlsCfg := &tls.Config{
		Certificates:          []tls.Certificate{cert},
		ClientAuth:            tls.RequireAnyClientCert,
		VerifyPeerCertificate: s.VerifyClientCertificateFunc,
	}

	for {
		select {
		case <-s.stopChan:
			return nil
		default:
			// Check listener.
		}

		// Reset deadline for the listener to stop blocking on accepting
		// connections and allow shutdown.
		if err := listener.SetDeadline(time.Now().Add(listenTimeout)); err != nil {
			log.Print(err.Error())
		}

		tcpConnection, err := listener.AcceptTCP()
		if err != nil {
			if opErr, ok := err.(net.Error); ok && opErr.Timeout() {
				// Not an error, the deadline timed out.
				continue
			}

			log.Printf("could not accept connection: %s", err.Error())
			continue
		}

		// The connection must be allowed to be opened for up to 10 minutes
		// without any manual activity so we enable keepalive on the TCP
		// socket.
		if err := tcpConnection.SetKeepAlive(true); err != nil {
			log.Printf("could not set keepalive")
			continue
		}

		if err := tcpConnection.SetKeepAlivePeriod(1 * time.Minute); err != nil {
			log.Printf("could set keepalive period")
			continue
		}

		// Wrap the TCP conn in a TLS server to enable handshake.
		tlsConnection := tls.Server(tcpConnection, tlsCfg)

		go s.processConnection(tlsConnection)
	}
}

// Stop will close the channel making no new regquests being processed and then
// drain all ongoing requests til they're done.
func (s *Server) Stop() {
	log.Print("stopping listener channel")
	close(s.stopChan)
}

func (s *Server) processConnection(conn *tls.Conn) {
	s.wg.Add(1)

	// Implement panic recovery and closing of the connection.
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("could not close connection: %s", err)
		}

		if e := recover(); e != nil {
			log.Printf("error %s", e)
		}

		s.wg.Done()
	}()

	// Check if we've got proxy protocol which we should have if behind load
	// balancer. The proxyproto package will return the address of the socket
	// if no proxy protocol is found.
	proxyConn := proxyproto.NewConn(conn, requestReadTimeout)
	remoteAddress, ok := proxyConn.RemoteAddr().(*net.TCPAddr)
	if !ok {
		log.Printf("not get peer IP")
		return
	}

	log.Printf("connection from %s", remoteAddress.IP)

	// Ensure we've finalized this handshake instead of perofrming it on the
	// first read/write.
	if err := conn.Handshake(); err != nil {
		log.Print("could not handshake")
		return
	}

	// Send greeting to client so they can verify the server, RFC 5734 section
	// 8, paragraph 5: http://www.rfc-editor.org/rfc/rfc5734.txt
	if err := Write(conn, s.GreetResponse()); err != nil {
		log.Print(err.Error())
		return
	}

	for {
		buffer, err := Read(conn)
		if err != nil {
			log.Print(err.Error())
			break
		}

		eppResult := s.handleCommand(buffer)
		eppResult.AddResultTag()

		if eppResult.Code.IsBye() {
			break
		}

		// TODO
		// break if conn.Disconnected
		// break if new connection > 4 found (max 4 sessions)
		// brak on SIGIN

		if err := Write(conn, eppResult.Response); err != nil {
			log.Println(err.Error())
		}

		break
	}
}

// Read will read from the connection by first parsing the content length
// header (read from the first 4 bytes) and read sent number of bytes from the
// connection.
func Read(conn *tls.Conn) ([]byte, error) {
	l := make([]byte, contentLengthHeaderBytes)
	reader := bufio.NewReader(conn)

	if _, err := reader.Read(l); err != nil {
		return nil, errors.Wrap(err, "could not read content length to connection")
	}

	buffer := binary.BigEndian.Uint32(l)
	dataBuffer := make([]byte, int(buffer)-contentLengthHeaderBytes)

	if _, err := reader.Read(dataBuffer); err != nil {
		return nil, errors.Wrap(err, "could not read XML to connection")
	}

	return dataBuffer, nil
}

// Write passed data to passed connection by first writing the content length
// in the first four bytes.
func Write(conn *tls.Conn, data interface{}) error {
	// Create a buffer, write the default XML header and encode the data with
	// an indent of 2 spaces and preserved newlines.
	xBuf := bytes.Buffer{}
	xBuf.WriteString(xml.Header)

	enc := xml.NewEncoder(&xBuf)
	enc.Indent("", "  ")

	if err := enc.EncodeElement(data, xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "epp",
		},
		Attr: []xml.Attr{
			{
				Name: xml.Name{
					Space: "xsi",
					Local: "schemaLocation",
				},
				Value: "urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd",
			},
			{
				Name: xml.Name{
					Space: "",
					Local: "xmlns",
				},
				Value: "urn:ietf:params:xml:ns:epp-1.0",
			},
			{
				Name: xml.Name{
					Space: "",
					Local: "xmlns:xsi",
				},
				Value: "http://www.w3.org/2001/XMLSchema-instance",
			},
		},
	}); err != nil {
		return err
	}

	l := make([]byte, contentLengthHeaderBytes)
	binary.BigEndian.PutUint32(l[0:], uint32(xBuf.Len()+contentLengthHeaderBytes))

	if _, err := conn.Write(l); err != nil {
		return errors.Wrap(err, "could not write content length to connection")
	}

	if _, err := conn.Write(xBuf.Bytes()); err != nil {
		return errors.Wrap(err, "could not write XML to connection")
	}

	return nil
}
