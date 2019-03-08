package eppserver

import (
	"strings"

	"aqwari.net/xml/xmltree"
	"github.com/pkg/errors"
)

const nsEPP = "urn:ietf:params:xml:ns:epp-1.0"

// Mux can be used to route different EPP messages to different handlers,
// depending on the message content.
//
//  m := Mux{}
//
//  m.AddNamespaceAlias("urn:ietf:params:xml:ns:domain-1.0", "domain")
//
//  m.AddHandler("hello", handleHello)
//  m.AddHandler("command/login", handleLogin)
//  m.AddHandler("command/check/urn:ietf:params:xml:ns:contact-1.0", handleCheckContact)
//  m.AddHandler("command/check/domain", handleCheckDomain)
type Mux struct {
	handlers         map[string]HandlerFunc
	namespaceAliases map[string]string
}

// NewMux will create and return a new Mux.
func NewMux() *Mux {
	m := &Mux{
		namespaceAliases: map[string]string{
			"urn:ietf:params:xml:ns:domain-1.0":  "domain",
			"urn:ietf:params:xml:ns:host-1.0":    "host",
			"urn:ietf:params:xml:ns:contact-1.0": "contact",
		},
	}

	return m
}

// AddNamespaceAlias will add an alias for the specified namespace. After the
// alias is added it can be used in routing. Multiple namespaces can be added
// to the same alias.
//  m.AddNamespaceAlias("urn:ietf:params:xml:ns:contact-1.0", "host-and-contact")
//  m.AddNamespaceAlias("urn:ietf:params:xml:ns:host-1.0", "host-and-contact")
func (m *Mux) AddNamespaceAlias(ns, alias string) {
	m.namespaceAliases[ns] = alias
}

// AddHandler will add a handler for the specified route.
// Routes are defined almost like xpath.
func (m *Mux) AddHandler(path string, handler HandlerFunc) {
	m.handlers[path] = handler
}

// Handle will handle an incoming message and route it to the correct handler.
// Pass the function to Server to use the Mux.
func (m *Mux) Handle(s *Session, d []byte) ([]byte, error) {
	root, err := xmltree.Parse(d)
	if err != nil {
		return nil, err
	}

	path, err := m.buildPath(root)
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

func (m *Mux) buildPath(root *xmltree.Element) (string, error) {
	// Ensure the start element is <epp>.
	if root.Name.Space != nsEPP || root.Name.Local != "epp" {
		return "", errors.New("missing <epp> tag")
	}

	// We should only have one element inside the <epp> tag.
	if len(root.Children) != 1 {
		return "", errors.New("<epp> should contain one element")
	}

	el := root.Children[0]
	if el.Name.Local != "command" {
		// It's not a command, so we'll just use this tag as the route.
		return el.Name.Local, nil
	}

	pathParts := []string{"command"}
	for _, child := range el.Children {
		name := child.Name.Local

		switch name {
		case "extension", "clTRID":
			// These tags can exist in a command but are always available
			// so we won't do any routing based on them.
			continue
		}

		switch name {
		case "login", "logout", "poll":
			// Login, Logout and Poll are commands defined in eppcom-1.0.xml
			// and does not need any further routing.
			pathParts = append(pathParts, name)
		default:
			// Other commands can be executed on multiple types of objects
			// that are defined by their namespace.
			ns := child.Children[0].Name.Space

			if alias, ok := m.namespaceAliases[ns]; ok {
				ns = alias
			}

			pathParts = append(pathParts, name, ns)
		}

		break
	}

	return strings.Join(pathParts, "/"), nil
}
