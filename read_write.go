package epp

import (
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"aqwari.net/xml/xmltree"
	"github.com/bombsimon/epp-go/types"
)

const (
	rootLocalName = "epp"
)

// ReadMessage reads one full message from r.
func ReadMessage(conn net.Conn) ([]byte, error) {
	// https://tools.ietf.org/html/rfc5734#section-4
	var totalSize uint32

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
				Local: "xmlns:xsi",
			},
			Value: "http://www.w3.org/2001/XMLSchema-instance",
		},
		{
			Name: xml.Name{
				Space: "",
				Local: "xsi:schemaLocation",
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
			Value: types.NameSpaceEPP10,
		},
	}
}

// Encode will take a type that can be marshalled to XML, add the EPP staring
// tag for all registered namespaces and return the XML as a byte slice.
func Encode(data interface{}, xmlAttributes []xml.Attr) ([]byte, error) {
	// Marshal input data to XML, assume types implement required marshaling
	// tags and features.
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	document, err := xmltree.Parse(b)
	if err != nil {
		return nil, err
	}

	addNameSpaceAlias(document, false)

	// Replace the document root element with a proper EPP tag.
	document.StartElement = xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: rootLocalName,
		},
		Attr: xmlAttributes,
	}

	// Marshal the xmltree after fixing name spaces and attributes.
	xmlBytes := xmltree.MarshalIndent(document, "", "  ")

	// Add XML header to the marshalled document.
	xmlBytes = append([]byte(xml.Header), xmlBytes...)

	return xmlBytes, nil
}

// addNameSpaceAlias will check each node/element in the XML tree and if the
// node has an xml.Name.Space value set an alias will be created and then added
// to all child nodes. The alias will only be setup for the root element.
func addNameSpaceAlias(document *xmltree.Element, nsAdded bool) *xmltree.Element {
	namespaceAliases := map[string]string{
		types.NameSpaceDomain:   "domain",
		types.NameSpaceHost:     "host",
		types.NameSpaceContact:  "contact",
		types.NameSpaceDNSSEC10: "sed",
		types.NameSpaceDNSSEC11: "sec",
		types.NameSpaceIIS12:    "iis",
	}

	if document.Name.Space != "" {
		alias, ok := namespaceAliases[document.Name.Space]
		if !ok {
			return nil
		}

		if !nsAdded {
			xmlns := fmt.Sprintf("xmlns:%s", alias)
			document.SetAttr("", xmlns, document.Name.Space)

			// Namespace alias is now added so flip to true to skip child
			// elements.
			nsAdded = true
		}

		document.Name.Local = fmt.Sprintf("%s:%s", alias, document.Name.Local)
	}

	for i, child := range document.Children {
		document.Children[i] = *addNameSpaceAlias(&child, nsAdded)
	}

	return document
}
