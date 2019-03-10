package epp

import (
	"time"

	ws "github.com/bombsimon/epp-server/types"
)

// GreetResponse generates the greet command.
func (s *Server) GreetResponse() ws.EppType {
	return ws.EppType{
		Greeting: &ws.GreetingType{
			SvID: ws.SIDType{
				Id: ws.ClIDType("asdb"),
			},
			SvDate: time.Now(),
		},
	}
}

// LoginResponse returns a successful login response.
func (s *Server) LoginResponse() ws.EppType {
	return ws.EppType{
		Response: &ws.ResponseType{
			TrID: ws.TrIDType{
				ClTRID: ws.TrIDStringType("XXX-99999"),
				SvTRID: ws.TrIDStringType("044"),
			},
		},
	}
}
