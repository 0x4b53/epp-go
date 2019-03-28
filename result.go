package epp

import (
	"fmt"

	"github.com/bombsimon/epp-go/types"
)

// ResultCode represents a result code from the EPP server.
type ResultCode int

// EPP constant types represents EPP result codes. For reference, see RFC5730
// section 3, "Result Codes".
const (
	EppOk                         ResultCode = 1000
	EppOkPending                  ResultCode = 1001
	EppOkNoMessages               ResultCode = 1300
	EppOkMessages                 ResultCode = 1301
	EppOkBye                      ResultCode = 1500
	EppUnknownCommand             ResultCode = 2000
	EppSyntaxError                ResultCode = 2001
	EppUseError                   ResultCode = 2002
	EppMissingParam               ResultCode = 2003
	EppParamRangeError            ResultCode = 2004
	EppParamSyntaxError           ResultCode = 2005
	EppUnimplementedVersion       ResultCode = 2100
	EppUnimplementedCommand       ResultCode = 2101
	EppUnimplementedOption        ResultCode = 2102
	EppUnimplementedExtension     ResultCode = 2103
	EppBillingFailure             ResultCode = 2104
	EppNotRenewable               ResultCode = 2105
	EppNotTransferrable           ResultCode = 2106
	EppAuthenticationError        ResultCode = 2200
	EppAuthorisationError         ResultCode = 2201
	EppInvalidAuthInfo            ResultCode = 2202
	EppObjectPendingTransfer      ResultCode = 2300
	EppObjectNotPendingTransfer   ResultCode = 2301
	EppObjectExists               ResultCode = 2302
	EppObjectDoesNotExist         ResultCode = 2303
	EppStatusProhibitsOp          ResultCode = 2304
	EppAssocProhibitsOp           ResultCode = 2305
	EppParamPolicyError           ResultCode = 2306
	EppUnimplementedObjectService ResultCode = 2307
	EppDataMgmtPolicyViolation    ResultCode = 2308
	EppCommandFailed              ResultCode = 2400
	EppCommandFailedBye           ResultCode = 2500
	EppAuthFailedBye              ResultCode = 2501
	EppSessionLimitExceededBye    ResultCode = 2502
)

// Code returns the integer code for the result code.
func (rs ResultCode) Code() int {
	return int(rs)
}

// Message returns the message to be embedded along the code.
func (rs ResultCode) Message() string {
	switch rs {
	case EppOk:
		return "Command completed successfully"
	case EppOkPending:
		return "Command completed successfully; action pending"
	case EppOkNoMessages:
		return "Command completed successfully; no messages"
	case EppOkMessages:
		return "Command completed successfully; ack to dequeue"
	case EppOkBye:
		return "Command completed successfully; ending session"
	case EppUnknownCommand:
		return "Unknown command"
	case EppSyntaxError:
		return "Command syntax error"
	case EppUseError:
		return "Command use error"
	case EppMissingParam:
		return "Required parameter missing"
	case EppParamRangeError:
		return "Parameter value range error"
	case EppParamSyntaxError:
		return "Parameter value syntax error"
	case EppUnimplementedVersion:
		return "Unimplemented protocol version"
	case EppUnimplementedCommand:
		return "Unimplemented command"
	case EppUnimplementedOption:
		return "Unimplemented option"
	case EppUnimplementedExtension:
		return "Unimplemented extension"
	case EppBillingFailure:
		return "Billing failure"
	case EppNotRenewable:
		return "Object is not eligible for renewal"
	case EppNotTransferrable:
		return "Object is not eligible for transfer"
	case EppAuthenticationError:
		return "Authentication error"
	case EppAuthorisationError:
		return "Authorization error"
	case EppInvalidAuthInfo:
		return "Invalid authorization information"
	case EppObjectPendingTransfer:
		return "Object pending transfer"
	case EppObjectNotPendingTransfer:
		return "Object not pending transfer"
	case EppObjectExists:
		return "Object exists"
	case EppObjectDoesNotExist:
		return "Object does not exist"
	case EppStatusProhibitsOp:
		return "Object status prohibits operation"
	case EppAssocProhibitsOp:
		return "Object association prohibits operation"
	case EppParamPolicyError:
		return "Parameter value policy error"
	case EppUnimplementedObjectService:
		return "Unimplemented object service"
	case EppDataMgmtPolicyViolation:
		return "Data management policy violation"
	case EppCommandFailed:
		return "Command failed"
	case EppCommandFailedBye:
		return "Command failed; server closing connection"
	case EppAuthFailedBye:
		return "Authentication error; server closing connection"
	case EppSessionLimitExceededBye:
		return "Session limit exceeded; server closing connection"
	default:
		return fmt.Sprintf("Code was %d", rs)
	}
}

// IsBye returns true if the result code is a connection management result code
// which should terminate the connection.
func (rs ResultCode) IsBye() bool {
	switch rs {
	case
		EppOkBye,
		EppCommandFailedBye,
		EppAuthFailedBye,
		EppSessionLimitExceededBye:
		return true
	default:
		return false
	}
}

// CreateErrorResponse will create a response with a given code, message and value
// which may be marshalled to XML and pass to WriteMessage to write a proper EPP
// response to the socket.
func CreateErrorResponse(code ResultCode, reason string) types.Response {
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
