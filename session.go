package epp

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	xsd "github.com/lestrrat-go/libxml2/xsd"
	uuid "github.com/satori/go.uuid"
)

// HandlerFunc represents a function for an EPP message.
type HandlerFunc func(*Session, []byte) ([]byte, error)

// GreetFunc represents a function handling a greeting for the EPP server.
type GreetFunc func(*Session) ([]byte, error)

// Session is an active connection to the EPP server.
type Session struct {
	// ConnectionState holds the state of the TLS connection initiated while
	// doing the handshake with the server.
	ConnectionState func() tls.ConnectionState

	// IdleTimeout is the maximum timeout to idle on the server before being
	// disconnected. Each time traffic is sent to the server the idle ticker is
	// reset and a new duration of IdleTimeout is allowed.
	IdleTimeout time.Duration

	// SessionID is a unique ID to use to identify a specific session.
	SessionID string

	// SessionTimeout is the max duration allowed for a single session. If a
	// client has been connected when this limit is reach it will be
	// disconnedted after the current command being processed has finished.
	SessionTimeout time.Duration

	// conn holds the TCP connection with a client.
	conn net.Conn

	// greeting holds the function that will generate the XML printed while
	// greeting clients connection to the server.
	greeting GreetFunc

	// handler holds a function that will receive each request to the server.
	handler HandlerFunc

	// stopChan is used to tell the session to terminate.
	stopChan chan struct{}

	// validator is a type that implements the validator interface. The
	// validator interface should be able to validate XML against an XSD schema
	// (or any other way). If the validator is a non nil value all incomming
	// *and* outgoing data will be passed through the validator. Type
	// implementing this interface using libxml2 bindings is available in the
	// library.
	validator Validator
}

// NewSession will create a new Session.
func NewSession(conn *tls.Conn, handler HandlerFunc, greeting GreetFunc, idleTimeout, sessionTimeout time.Duration, validator Validator) *Session {
	sessionID := uuid.Must(uuid.NewV4()).String()

	s := &Session{
		SessionID:       sessionID,
		SessionTimeout:  sessionTimeout,
		IdleTimeout:     idleTimeout,
		ConnectionState: conn.ConnectionState,
		conn:            conn,
		greeting:        greeting,
		handler:         handler,
		stopChan:        make(chan struct{}),
		validator:       validator,
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
		s.validator.Free()
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
