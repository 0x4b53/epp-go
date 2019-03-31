package epp

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Server represents the server handling requests.
type Server struct {
	IdleTimeout      time.Duration
	MaxSessionLength time.Duration
	TLSConfig        *tls.Config
	Handler          HandlerFunc
	Greeting         GreetFunc
	Addr             string

	// Sessions will contain all the currently active sessions.
	Sessions   map[string]*Session
	sessionsMu sync.Mutex
	sessionsWg sync.WaitGroup

	onStarteds []func()
	stopChan   chan struct{}
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
		tlsConn.ConnectionState,
		uuid.Must(uuid.NewV4()).String(),
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
