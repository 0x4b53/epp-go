package eppserver

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"net"
	"time"
)

// ReadMessage reads one full message from r.
func ReadMessage(conn net.Conn) ([]byte, error) {
	// https://tools.ietf.org/html/rfc5734#section-4
	var totalSize uint32 = 0
	err := binary.Read(conn, binary.BigEndian, &totalSize)
	if err != nil {
		return nil, err
	}

	headerSize := binary.Size(totalSize)
	contentSize := int(totalSize) - headerSize

	// Ensure a reasonable time for reading the message.
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, contentSize)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// WriteMessage writes data to w with the correct header.
func WriteMessage(conn net.Conn, data []byte) error {
	// Begin by writing the len(b) as Big Endian uint32, including the
	// size of the content length header.
	// https://tools.ietf.org/html/rfc5734#section-4
	contentSize := len(data)
	headerSize := binary.Size(uint32(contentSize))
	totalSize := contentSize + headerSize

	// Bounds check.
	if totalSize > math.MaxUint32 {
		return errors.New("content is too large")
	}

	err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}

	err = binary.Write(conn, binary.BigEndian, uint32(totalSize))
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
