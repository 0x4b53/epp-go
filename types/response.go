package types

import "time"

// Response represents an EPP response.
type Response struct {
	Result        []Result      `xml:"response>result"`
	MessageQ      *MessageQueue `xml:"response>msgQ,omitempty"`
	ResultData    interface{}   `xml:"response>resData,omitempty"`
	Extension     interface{}   `xml:"response>extension,omitempty"`
	TransactionID TransactionID `xml:"response>trID"`
}

// TransactionID represents transaction IDs for the client and the server.
type TransactionID struct {
	ClientTransactionID string `xml:"clTRID,omitempty"`
	ServerTransactionID string `xml:"svTRID"`
}

// MessageQueue represents a message queue for client retrieval.
type MessageQueue struct {
	QueueDate *time.Time `xml:"qDate,omitempty"`
	Message   string     `xml:"msg,omitempty"`
	Count     int        `xml:"count,attr"`
	ID        string     `xml:"id,attr"`
}

// Result represents the result in a EPP response.
type Result struct {
	Code          int                 `xml:"code,attr"`
	Message       string              `xml:"msg"`
	Value         interface{}         `xml:"value"`
	ExternalValue *ExternalErrorValue `xml:"extValue,omitempty"`
}

// ExternalErrorValue represents the response in the extValeu tag.
type ExternalErrorValue struct {
	Value  interface{} `xml:"value"`
	Reason string      `xml:"reason"`
}
