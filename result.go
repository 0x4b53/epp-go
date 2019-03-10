package epp

import (
	"fmt"

	ws "github.com/bombsimon/epp-server/types"
)

// EppResultCode represents a result code from the EPP server.
type EppResultCode int

// EPP constant types represents EPP result codes. For reference, see RFC5730
// section 3, "Result Codes".
const (
	EppOk                         EppResultCode = 1000
	EppOkPending                  EppResultCode = 1001
	EppOkNoMessages               EppResultCode = 1300
	EppOkMessages                 EppResultCode = 1301
	EppOkBye                      EppResultCode = 1500
	EppUnknownCommand             EppResultCode = 2000
	EppSyntaxError                EppResultCode = 2001
	EppUseError                   EppResultCode = 2002
	EppMissingParam               EppResultCode = 2003
	EppParamRangeError            EppResultCode = 2004
	EppParamSyntaxError           EppResultCode = 2005
	EppUnimplementedVersion       EppResultCode = 2100
	EppUnimplementedCommand       EppResultCode = 2101
	EppUnimplementedOption        EppResultCode = 2102
	EppUnimplementedExtension     EppResultCode = 2103
	EppBillingFailure             EppResultCode = 2104
	EppNotRenewable               EppResultCode = 2105
	EppNotTransferrable           EppResultCode = 2106
	EppAuthenticationError        EppResultCode = 2200
	EppAuthorisationError         EppResultCode = 2201
	EppAuthorizationError         EppResultCode = 2201
	EppInvalidAuthInfo            EppResultCode = 2202
	EppObjectPendingTransfer      EppResultCode = 2300
	EppObjectNotPendingTransfer   EppResultCode = 2301
	EppObjectExists               EppResultCode = 2302
	EppObjectDoesNotExist         EppResultCode = 2303
	EppStatusProhibitsOp          EppResultCode = 2304
	EppAssocProhibitsOp           EppResultCode = 2305
	EppParamPolicyError           EppResultCode = 2306
	EppUnimplementedObjectService EppResultCode = 2307
	EppDataMgmtPolicyViolation    EppResultCode = 2308
	EppCommandFailed              EppResultCode = 2400
	EppCommandFailedBye           EppResultCode = 2500
	EppAuthFailedBye              EppResultCode = 2501
	EppSessionLimitExceededBye    EppResultCode = 2502
)

// Message returns the message to be embedded along the code.
func (e EppResultCode) Message() string {
	switch e {
	case EppOk:
		return "Command completed successfully"
	default:
		return fmt.Sprintf("Code was %d", int(e))
	}
}

// IsBye returns true if the result code is a connection management result code
// which should terminate the connection.
func (e EppResultCode) IsBye() bool {
	switch e {
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

// EppResult represents the result after a query to the server has been
// processed.
type EppResult struct {
	Error    error
	Code     EppResultCode
	Response ws.EppType
}

// AddResultTag will add the <result> tag to the EPP response and include a
// message. If an error exists, the message from the error will be used. If
// not, the message from the EppResult will be used.
func (er *EppResult) AddResultTag() {
	message := er.Code.Message()
	if er.Error != nil {
		message = er.Error.Error()
	}

	// TODO: No response added?!
	if er.Response.Response == nil {
		return
	}

	er.Response.Response.Result = []ws.ResultType{
		{
			Msg: ws.MsgType{
				Value: message,
			},
			Code: ws.ResultCodeType(er.Code),
		},
	}
}
