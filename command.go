package eppserver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	ws "github.com/bombsimon/epp-server/types"
)

// XPath are constants for XPath strings to find commands.
const (
	XPathHello           = "/epp/hello"
	XPathCommandLogin    = "/epp/command/login"
	XPathCommandLogout   = "/epp/command/logout"
	XPathCommandPoll     = "/epp/command/poll"
	XPathCommandCheck    = "/epp/command/check"
	XPathCommandCreate   = "/epp/command/create"
	XPathCommandDelete   = "/epp/command/delete"
	XPathCommandInfo     = "/epp/command/info"
	XPathCommandRenew    = "/epp/command/renew"
	XPathCommandTransfer = "/epp/command/transfer"
	XPathCommandUpdate   = "/epp/command/update"
)

func (s *Server) handleCommand(buffer []byte) EppResult {
	dataBufferString := string(buffer)

	doc, err := xmlquery.Parse(strings.NewReader(dataBufferString))
	if err != nil {
		fmt.Println("!!!", err.Error())
	}

	/*
		Check non prefix/session commands:
		/epp/hello
		/epp/command/login
		/epp/command/logout
		/epp/command/poll

		The commands below has a prefix for object type, i.e.
		<command>
		  <check>
		    <domain:check>
		      <domain:name>sawert.se</domain:name>
		    </domain:check>
		  </check>
		</command>

		/epp/command/->
			/check
			/create
			/delete
			/info
			/renew
			/transfer
			/update
	*/

	if node := xmlquery.FindOne(doc, XPathHello); node != nil {
		return s.HelloCommand(node)
	}

	if node := xmlquery.FindOne(doc, XPathCommandLogin); node != nil {
		return s.LoginCommand(node)
	}

	return EppResult{
		Code:     EppSyntaxError,
		Response: ws.EppType{},
		Error:    errors.New("No valid command found"),
	}
}

// HelloCommand handles the hello command to the server.
func (s *Server) HelloCommand(node *xmlquery.Node) EppResult {
	result := EppResult{
		Code:     EppOk,
		Response: ws.EppType{},
	}

	// Only one node
	// Has no data
	// Has no attributes

	result.Response = s.GreetResponse()

	return result
}

// LoginCommand handles the login command to the server.
func (s *Server) LoginCommand(node *xmlquery.Node) EppResult {
	result := EppResult{
		Code:     EppOk,
		Response: ws.EppType{},
	}

	username := node.SelectElement("clID").InnerText()
	password := node.SelectElement("pw").InnerText()

	if username == "test" && password == "test" {
		result.Response = s.LoginResponse()

	} else {
		result.Code = EppAuthFailedBye
		result.Error = errors.New("bad login")
	}

	return result
}
