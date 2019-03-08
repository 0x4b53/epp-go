package eppserver

import (
	"crypto/tls"
	"net"
	"time"
)

type Client struct {
	TLSConfig *tls.Config
	messages  chan []byte
	conn      net.Conn
}

func (c *Client) Connect(server string) ([]byte, error) {
	if c.TLSConfig == nil {
		c.TLSConfig = &tls.Config{}
	}
	conn, err := tls.Dial("tcp", server, c.TLSConfig)
	if err != nil {
		return nil, err
	}

	// Read the greeting.
	greeting, err := ReadMessage(conn)
	if err != nil {
		return nil, err
	}

	c.conn = conn

	return greeting, nil
}

func (c *Client) Send(data []byte) ([]byte, error) {
	err := WriteMessage(c.conn, data)
	if err != nil {
		_ = c.conn.Close()
		return nil, err
	}

	_ = c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	msg, err := ReadMessage(c.conn)
	if err != nil {
		_ = c.conn.Close()
		return nil, err
	}

	return msg, nil
}
