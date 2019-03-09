package epp

import (
	"fmt"
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
	EppAuthorizationError         ResultCode = 2201
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

// Message returns the message to be embedded along the code.
func (e ResultCode) Message() string {
	switch e {
	case EppOk:
		return "Command completed successfully"
	default:
		return fmt.Sprintf("Code was %d", int(e))
	}
}

// IsBye returns true if the result code is a connection management result code
// which should terminate the connection.
func (e ResultCode) IsBye() bool {
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

// Result represents the result after a query to the server has been processed.
type Result struct {
	Error    error
	Code     ResultCode
	Response interface{}
}

// AddResultTag will add the <result> tag to the EPP response and include a
// message. If an error exists, the message from the error will be used. If
// not, the message from the Result will be used.
func (er *Result) AddResultTag() {
	// TODO
}
