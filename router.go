package eppserver

import (
	"strings"

	"aqwari.net/xml/xmltree"
	"github.com/pkg/errors"
)

const nsEPP = "urn:ietf:params:xml:ns:epp-1.0"

var routingNamespaces = map[string]string{
	"urn:ietf:params:xml:ns:domain-1.0":  "domain",
	"urn:ietf:params:xml:ns:host-1.0":    "host",
	"urn:ietf:params:xml:ns:contact-1.0": "contact",
}

type Mux struct {
	handlers map[string]HandlerFunc
}

func (m *Mux) AddHandler(p string, h HandlerFunc) {
	m.handlers[p] = h
}

func (m *Mux) Handle(s *Session, d []byte) ([]byte, error) {
	root, err := xmltree.Parse(d)
	if err != nil {
		return nil, err
	}

	path, err := buildPath(root)
	if err != nil {
		return nil, err
	}

	h, ok := m.handlers[path]
	if !ok {
		// TODO
		return nil, errors.Errorf("no handler for %s", path)
	}

	return h(s, d)
}

func buildPath(root *xmltree.Element) (string, error) {
	// 1. Ensure the start element is <epp>.
	if root.Name.Space != nsEPP || root.Name.Local != "epp" {
		return "", errors.New("missing <epp> tag")
	}

	// 2. We should only have one element inside the <epp> tag.
	if len(root.Children) != 1 {
		return "", errors.New("<epp> should contain one element")
	}
	el := root.Children[0]
	if el.Name.Local != "command" {
		return el.Name.Local, nil
	}

	pathParts := []string{"command"}
	for _, child := range el.Children {
		name := child.Name.Local

		switch name {
		case "extension", "clTRID":
			continue
		}

		switch name {
		case "login", "logout", "poll":
			pathParts = append(pathParts, name)
		default:
			pathParts = append(pathParts, name)

			namespace, ok := routingNamespaces[child.Children[0].Name.Space]
			if !ok {
				namespace = "unknown"
			}
			pathParts = append(pathParts, namespace)
		}

		break
	}

	return strings.Join(pathParts, ">"), nil
}
