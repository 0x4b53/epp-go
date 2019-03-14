package epp

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"io"
	"math"
	"net"
	"time"

	"github.com/bombsimon/epp-go/types"
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

// ServerXMLAttributes defines the default attributes from the server response.
func ServerXMLAttributes() []xml.Attr {
	return []xml.Attr{
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
				Local: "xsi",
			},
			Value: "http://www.w3.org/2001/XMLSchema-instance",
		},
		{
			Name: xml.Name{
				Space: "xsi",
				Local: "schemaLocation",
			},
			Value: "urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd",
		},
	}
}

// ClientXMLAttributes defines the default attributes in the client request.
func ClientXMLAttributes() []xml.Attr {
	return []xml.Attr{
		{
			Name: xml.Name{
				Space: "",
				Local: "xmlns",
			},
			Value: "urn:ietf:params:xml:ns:epp-1.0",
		},
	}
}

// Encode will take a type that can be marshalled to XML, add the EPP staring
// tag for all registered namespaces and return the XML as a byte slice.
func Encode(data interface{}, xmlAttributes []xml.Attr) ([]byte, error) {
	// Create a buffer, write the default XML header and encode the data with
	// an indent of 2 spaces and preserved newlines.
	buf := bytes.Buffer{}
	buf.WriteString(xml.Header)

	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	err := enc.EncodeElement(data, xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "epp",
		},
		Attr: xmlAttributes,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// CreateResponse will create a response with a given code, message and value
// which may be marshalled to XML and pass to WriteMessage to write a proper EPP
// response to the socket.
func CreateResponse(code ResultCode, reason string) types.Response {
	return types.Response{
		Result: []types.Result{
			{
				Code:    code.Code(),
				Message: code.Message(),
				ExternalValue: types.ExternalErrorValue{
					Reason: reason,
				},
			},
		},
	}
}
