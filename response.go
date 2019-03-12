package epp

import (
	"bytes"
	"encoding/xml"
	"time"

	"github.com/bombsimon/epp-go/types"
)

// GreetResponse generates the greet command.
func (s *Server) GreetResponse() types.EPPGreeting {
	return types.EPPGreeting{
		Greeting: types.Greeting{
			ServerID:   "default-server",
			ServerDate: time.Now(),
			ServiceMenu: types.ServiceMenu{
				Version: []string{"1.0"},
			},
		},
	}
}

// Encode will take a type that can be marshalled to XML, add the EPP staring
// tag for all registered namespaces and return the XML as a byte slice.
func (s *Server) Encode(data interface{}, ns map[string]string) ([]byte, error) {
	// Create a buffer, write the default XML header and encode the data with
	// an indent of 2 spaces and preserved newlines.
	buf := bytes.Buffer{}
	buf.WriteString(xml.Header)

	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	registeredNamespaces := []xml.Attr{
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

	for value, local := range ns {
		registeredNamespaces = append(registeredNamespaces, xml.Attr{
			Name: xml.Name{
				Space: "xmlns",
				Local: local,
			},
			Value: value,
		})
	}

	err := enc.EncodeElement(data, xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "epp",
		},
		Attr: registeredNamespaces,
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
