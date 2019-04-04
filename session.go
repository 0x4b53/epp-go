package epp

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	xsd "github.com/lestrrat-go/libxml2/xsd"
)

// HandlerFunc represents a function for an EPP message.
type HandlerFunc func(*Session, []byte) ([]byte, error)

// GreetFunc represents a function handling a greeting for the EPP server.
type GreetFunc func(*Session) ([]byte, error)

// Session is an active connection to the EPP server.
type Session struct {
	stopChan  chan struct{}
	conn      net.Conn
	handler   HandlerFunc
	greeting  GreetFunc
	validator *Validator

	Data            map[string]interface{}
	SessionID       string
	SessionTimeout  time.Duration
	IdleTimeout     time.Duration
	ConnectionState func() tls.ConnectionState
}

// NewSession will create a new Session.
func NewSession(conn net.Conn, handler HandlerFunc, greeting GreetFunc, tlsStateFunc func() tls.ConnectionState, sessionID string) *Session {
	validator, err := NewValidator("xml/index.xsd")
	if err != nil {
		panic(err)
	}

	s := &Session{
		stopChan:        make(chan struct{}),
		conn:            conn,
		handler:         handler,
		greeting:        greeting,
		validator:       validator,
		Data:            map[string]interface{}{},
		SessionID:       sessionID,
		SessionTimeout:  1 * time.Hour,
		IdleTimeout:     10 * time.Minute,
		ConnectionState: tlsStateFunc,
	}

	return s
}

// run will start the session.
func (s *Session) run() error {
	defer s.conn.Close()

	response, err := s.greeting(s)
	if err != nil {
		// TODO: Write response code and message?
		return err
	}

	if err := s.validate(response); err != nil {
		return err
	}

	err = WriteMessage(s.conn, response)
	if err != nil {
		return err
	}

	sessionTimeout := time.After(s.SessionTimeout)
	idleTimeout := time.After(s.IdleTimeout)

	for {
		select {
		case <-s.stopChan:
			log.Printf("stopping server, ending session %s", s.SessionID)

			return nil
		case <-sessionTimeout:
			log.Printf("session has been active for 1 hour, ending session %s", s.SessionID)

			return nil
		case <-idleTimeout:
			log.Printf("session has been idle for 10 minutes, ending session %s", s.SessionID)

			return nil
		default:
			// Go on...
		}

		err = s.conn.SetDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			return err
		}

		message, err := ReadMessage(s.conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			return err
		}

		log.Printf("handling incomming message")

		if err := s.validate(message); err != nil {
			return err
		}

		// Handle Message:
		response, err = s.handler(s, message)
		if err != nil {
			return err
		}

		if err := s.validate(response); err != nil {
			return err
		}

		err = WriteMessage(s.conn, response)
		if err != nil {
			return err
		}

		// Extend the idle timeout.
		idleTimeout = time.After(s.IdleTimeout)
	}
}

// Close will tell the session to close.
func (s *Session) Close() error {
	close(s.stopChan)

	if s.validator != nil {
		s.validator.Schema.Free()
	}

	return nil
}

func (s *Session) validate(data []byte) error {
	if s.validator == nil {
		return nil
	}

	if err := s.validator.Validate(data); err != nil {
		if xErr, ok := err.(xsd.SchemaValidationError); ok {
			for _, e := range xErr.Errors() {
				log.Printf("error: %s", e.Error())
			}
		}

		return err
	}

	return nil
}
