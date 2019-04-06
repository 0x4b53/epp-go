package epp

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Server represents the server handling requests.
type Server struct {
	// Addr is the address to use when listening to incomming TCP connections.
	// This should be set to ':700' to access incomming traffic on any interface
	// for the default EPP port 700.
	Addr string

	// Greeting is the function to exdecute while perofrming the greeting. This
	// function should return the exact XML data that should be written to the
	// socket. If an error is return, no greeting will be presented ant the
	// socket will be closed.
	Greeting GreetFunc

	// Handler is the function that will receive the current session and the
	// data printed to the socket for each request to the server. Use an epp mux
	// from this library to route messages based on it's content.
	Handler HandlerFunc

	// IdleTimeout is the maximum timeout to idle on the server before being
	// disconnected. Each time traffic is sent to the server the idle ticker is
	// reset and a new duration of IdleTimeout is allowed.
	IdleTimeout time.Duration

	// SessionTimeout is the max duration allowed for a single session. If a
	// client has been connected when this limit is reach it will be
	// disconnedted after the current command being processed has finished.
	SessionTimeout time.Duration

	// TLSConfig is the server TLS config with configuration such as
	// certificates, client auth etcetera.
	TLSConfig *tls.Config

	// Validator is a type that implements the validator interface. The
	// validator interface should be able to validate XML against an XSD schema
	// (or any other way). If the validator is a non nil value all incomming
	// *and* outgoing data will be passed through the validator. Type
	// implementing this interface using libxml2 bindings is available in the
	// library.
	Validator Validator

	// Sessions will contain all the currently active sessions.
	Sessions map[string]*Session

	// sessionMu is a mutex to use while reading and writing to the Sessions
	// liste to ensure thread safe access.
	sessionsMu sync.Mutex

	// sessionWg is a wait group used to ensure all ongoing sessions are
	// finished before closing the server.
	sessionsWg sync.WaitGroup

	// onStarted holds a list of functions that will be executed after the
	// server has been started.
	onStarteds []func()

	// stopChan is the channel that will be closed to tell when the server
	// should do a graceful shutdown.
	stopChan chan struct{}
}

// ListenAndServe will start the epp server.
func (s *Server) ListenAndServe() error {
	addr, err := net.ResolveTCPAddr("tcp", s.Addr)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	err = s.Serve(l)
	if err != nil {
		return err
	}

	return nil
}

// Serve will serve connections by listening on l.
func (s *Server) Serve(l *net.TCPListener) error {
	tlsConfig := &tls.Config{}

	s.sessionsWg = sync.WaitGroup{}
	s.stopChan = make(chan struct{})
	s.Sessions = map[string]*Session{}

	defer func() {
		if closeErr := l.Close(); closeErr != nil {
			fmt.Println(closeErr.Error())
		}
		s.sessionsWg.Wait()
	}()

	if s.TLSConfig != nil {
		tlsConfig = s.TLSConfig.Clone()
	}

	for _, f := range s.onStarteds {
		f()
	}

	var err error

	for {
		// Reset deadline for the listener to stop blocking on accepting
		// connections and allow shutdown.
		err = l.SetDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			return err
		}

		conn, err := l.AcceptTCP()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				select {
				case <-s.stopChan:
					return nil
				default:
					continue
				}
			}

			return err
		}

		// The connection must be allowed to be opened for up to 10 minutes
		// without any manual activity so we enable keepalive on the TCP
		// socket.
		if err := conn.SetKeepAlive(true); err != nil {
			log.Println(err.Error())
			continue
		}

		if err := conn.SetKeepAlivePeriod(1 * time.Minute); err != nil {
			log.Printf(err.Error())
			continue
		}

		go s.startSession(conn, tlsConfig)
	}
}

func (s *Server) startSession(conn net.Conn, tlsConfig *tls.Config) {
	// Initialize tls.
	tlsConn := tls.Server(conn, tlsConfig)

	err := tlsConn.Handshake()
	if err != nil {
		log.Println(err.Error())

		return
	}

	session := NewSession(
		tlsConn,
		s.Handler,
		s.Greeting,
		s.IdleTimeout,
		s.SessionTimeout,
		s.Validator,
	)

	// Ensure the session is added to our index.
	s.sessionsWg.Add(1)
	s.sessionsMu.Lock()
	s.Sessions[session.SessionID] = session
	s.sessionsMu.Unlock()

	// Ensure the session is removed from our session index when this function
	// exits.
	defer func() {
		s.sessionsMu.Lock()
		if _, ok := s.Sessions[session.SessionID]; ok {
			delete(s.Sessions, session.SessionID)
		}

		s.sessionsMu.Unlock()

		s.sessionsWg.Done()
		log.Println("session completed")
	}()

	log.Println("starting session", session.SessionID)

	err = session.run()
	if err != nil {
		log.Println(err)
	}
}

// OnStarted will register a function that is called when the server has
// finished it's startup.
func (s *Server) OnStarted(f func()) {
	s.onStarteds = append(s.onStarteds, f)
}

// Stop will close the channel making no new regquests being processed and then
// drain all ongoing requests til they're done.
func (s *Server) Stop() {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	log.Print("stopping listener channel")
	close(s.stopChan)

	for _, session := range s.Sessions {
		err := session.Close()
		if err != nil {
			log.Println("error closing session:", err.Error())
		}
	}
}
