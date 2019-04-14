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

	// OnStarteds holds a list of functions that will be executed after the
	// server has been started.
	OnStarteds []func()

	// SessionConfig holds the configuration to use for eachsession created.
	SessionConfig SessionConfig

	// TLSConfig is the server TLS config with configuration such as
	// certificates, client auth etcetera.
	TLSConfig *tls.Config

	// Sessions will contain all the currently active sessions.
	Sessions map[string]*Session

	// sessionMu is a mutex to use while reading and writing to the Sessions
	// liste to ensure thread safe access.
	sessionsMu sync.Mutex

	// sessionWg is a wait group used to ensure all ongoing sessions are
	// finished before closing the server.
	sessionsWg sync.WaitGroup

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
	s.sessionsWg = sync.WaitGroup{}
	s.stopChan = make(chan struct{})
	s.Sessions = map[string]*Session{}

	defer func() {
		if closeErr := l.Close(); closeErr != nil {
			fmt.Println(closeErr.Error())
		}

		s.sessionsWg.Wait()
	}()

	tlsConfig := &tls.Config{}

	// Use the same TLS config for the session if used on the server.
	if s.TLSConfig != nil {
		tlsConfig = s.TLSConfig.Clone()
	}

	// Perform user defined functions to execute each time the server is
	// started.
	for _, f := range s.OnStarteds {
		f()
	}

	for {
		// Reset deadline for the listener to stop blocking on accepting
		// connections and allow shutdown.
		if err := l.SetDeadline(time.Now().Add(1 * time.Second)); err != nil {
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

	session := NewSession(tlsConn, s.SessionConfig)

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

	if err = session.run(); err != nil {
		log.Println(err)
	}
}

// Stop will close the channel making no new regquests being processed and then
// drain all ongoing requests til they're done.
func (s *Server) Stop() {
	s.sessionsMu.Lock()
	defer s.sessionsMu.Unlock()

	log.Print("stopping listener channel")

	close(s.stopChan)

	for _, session := range s.Sessions {
		if err := session.Close(); err != nil {
			log.Println("error closing session:", err.Error())
		}
	}
}
