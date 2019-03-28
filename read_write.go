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
	rootLocalName          = "epp"
	nsCommandTagStartDepth = 3
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
			Value: types.NameSpaceEPP10,
		},
	}
}

// Encode will take a type that can be marshalled to XML, add the EPP staring
// tag for all registered namespaces and return the XML as a byte slice.
func Encode(data interface{}, xmlAttributes []xml.Attr, ns string) ([]byte, error) {
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

	// Check if a name space was passed and if so add the alias to all child
	// tags.
	addNameSpace(document, ns)

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

// addNameSpace will check if a name space was passed and find the first tag
// where it should be applied.
// TODO: This should work on any depth and tag, including extensions. I guess
// we must support multiple name spaces and aliases as well.
func addNameSpace(document *xmltree.Element, alias string) {
	if alias == "" {
		return
	}

	childElement := document

	// For default commands we shuld find the first element where we want to
	// apply the name space alias. If we run out of children before the n:th tag
	// is found we do nothing.
	for range make([]int, nsCommandTagStartDepth) {
		if len(childElement.Children) != 1 {
			return
		}

		childElement = &childElement.Children[0]
	}

	xmlns := fmt.Sprintf("xmlns:%s", alias)
	xmlnsValue := types.AliasToNameSpace(alias)

	// Add the attribute xmlns with the appropreate value based on name space.
	// TODO: Now we just support domain, contact and host.
	childElement.SetAttr("", xmlns, xmlnsValue)

	// Rename the local tag with the alias prefix, i.e. 'create' >
	// 'domain:create'.
	// TODO: Why can we not set element.Name.Space?
	childElement.Name.Local = fmt.Sprintf("%s:%s", alias, childElement.Name.Local)

	// Add the name space to each tag recursively.
	addNameSpaceToChildren(childElement.Children, alias)
}

// addNameSpaceToChildren is a recursive function which will add the alias
// prefix to all children passed to the function.
func addNameSpaceToChildren(children []xmltree.Element, alias string) {
	for i, element := range children {
		element.Name.Local = fmt.Sprintf("%s:%s", alias, element.Name.Local)
		children[i] = element

		if len(element.Children) > 0 {
			addNameSpaceToChildren(element.Children, alias)
		}
	}
}
