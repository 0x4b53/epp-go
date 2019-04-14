package epp

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	xsd "github.com/lestrrat-go/libxml2/xsd"
)

// HandlerFunc represents a function for an EPP message.
type HandlerFunc func(*Session, []byte) ([]byte, error)

// GreetFunc represents a function handling a greeting for the EPP server.
type GreetFunc func(*Session) ([]byte, error)

// SessionConfig represent the configuration passed to each new session being
// created.
type SessionConfig struct {
	// IdleTimeout is the maximum timeout to idle on the server before being
	// disconnected. Each time traffic is sent to the server the idle ticker is
	// reset and a new duration of IdleTimeout is allowed.
	IdleTimeout time.Duration

	// SessionTimeout is the max duration allowed for a single session. If a
	// client has been connected when this limit is reach it will be
	// disconnedted after the current command being processed has finished.
	SessionTimeout time.Duration

	// greeting holds the function that will generate the XML printed while
	// greeting clients connection to the server.
	Greeting GreetFunc

	// handler holds a function that will receive each request to the server.
	Handler HandlerFunc

	// validator is a type that implements the validator interface. The
	// validator interface should be able to validate XML against an XSD schema
	// (or any other way). If the validator is a non nil value all incomming
	// *and* outgoing data will be passed through the validator. Type
	// implementing this interface using libxml2 bindings is available in the
	// library.
	Validator Validator

	// OnCommands is a list of functions that will be executed on each command.
	// This is the place to put external code to handle after each command.
	OnCommands []func(sess *Session)
}

// Session is an active connection to the EPP server.
type Session struct {
	// ConnectionState holds the state of the TLS connection initiated while
	// doing the handshake with the server.
	ConnectionState func() tls.ConnectionState

	// SessionID is a unique ID to use to identify a specific session.
	SessionID string

	// conn holds the TCP connection with a client.
	conn net.Conn

	// stopChan is used to tell the session to terminate.
	stopChan chan struct{}

	// Se configurables details in SessionConfig
	IdleTimeout    time.Duration
	SessionTimeout time.Duration
	greeting       GreetFunc
	handler        HandlerFunc
	onCommands     []func(sess *Session)
	validator      Validator
}

// NewSession will create a new Session.
func NewSession(conn *tls.Conn, cfg SessionConfig) *Session {
	sessionID := uuid.New().String()

	s := &Session{
		SessionID:       sessionID,
		ConnectionState: conn.ConnectionState,
		conn:            conn,
		stopChan:        make(chan struct{}),
		IdleTimeout:     cfg.IdleTimeout,
		SessionTimeout:  cfg.SessionTimeout,
		greeting:        cfg.Greeting,
		handler:         cfg.Handler,
		onCommands:      cfg.OnCommands,
		validator:       cfg.Validator,
	}

	return s
}

// run will start the session.
func (s *Session) run() error {
	defer s.conn.Close()

	// Send the greeting to the client to do a proper greeting process, RFC5730,
	// 2.4
	response, err := s.greeting(s)
	if err != nil {
		// TODO: Write response code and message?
		return err
	}

	// Before the greeting is returned, ensure it's valid according to the EPP
	// XSD.
	if err := s.validate(response); err != nil {
		return err
	}

	// Write the greeting on the socket.
	err = WriteMessage(s.conn, response)
	if err != nil {
		return err
	}

	// Start the timers for the session and for idling.
	sessionTimeout := time.After(s.SessionTimeout)
	idleTimeout := time.After(s.IdleTimeout)

	for {
		select {
		case <-s.stopChan:
			log.Printf("stopping server, ending session %s", s.SessionID)

			return nil
		case <-sessionTimeout:
			log.Printf("session has been active for %v minutes, ending session %s", s.SessionTimeout.Minutes(), s.SessionID)

			return nil
		case <-idleTimeout:
			log.Printf("session has been idle for %v minutes, ending session %s", s.IdleTimeout.Minutes(), s.SessionID)

			return nil
		default:
			// We're not told to stop and we haven't reach the timeout for the
			// session or idling so just try to read from the socket again.
		}

		// Set the deadlien before reading to one second so we can check for
		// timeouts and stops each second to do proper timeout handling.
		err = s.conn.SetDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			return err
		}

		// Read from the socket and ensure that if we get an error we ignore it
		// if there were no activity on the socket.
		message, err := ReadMessage(s.conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			return err
		}

		// Before we execute the command, perform all functions defined to be
		// executed on each command. This might be where the user added rate
		// limiting or other things to handle before executing the command.
		for _, f := range s.onCommands {
			f(s)
		}

		// Validate the incomming XML data towards the RFC XSD.
		if err := s.validate(message); err != nil {
			return err
		}

		// Handle the message by passing it to the handler which may then take
		// action or route the message.
		response, err = s.handler(s, message)
		if err != nil {
			return err
		}

		// Validate the response to from the handler towards the XSD so we
		// don'tsend invalid XML to the client.
		if err := s.validate(response); err != nil {
			return err
		}

		// Write the message on the socket.
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
