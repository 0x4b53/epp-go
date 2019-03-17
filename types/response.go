package types

import "time"

// Response represents an EPP response.
type Response struct {
	Result        []Result      `xml:"response>result"`
	MessageQ      MessageQueue  `xml:"response>msgQ"`
	ResultData    interface{}   `xml:"response>resData"`
	Extension     interface{}   `xml:"response>extension"`
	TransactionID TransactionID `xml:"response>trID"`
}

// TransactionID ...
type TransactionID struct {
	ClientTransactionID string `xml:"clTRID"`
	ServerTransactionID string `xml:"svTRID"`
}

// MessageQueue ...
type MessageQueue struct {
	QueueDate time.Time `xml:"qDate"`
	Message   string    `xml:"msg"`
	Count     int       `xml:"count,attr"`
	ID        string    `xml:"id,attr"`
}

// Result represents the result in a EPP response.
type Result struct {
	Code          int                `xml:"code,attr"`
	Message       string             `xml:"msg"`
	Value         interface{}        `xml:"value"`
	ExternalValue ExternalErrorValue `xml:"extValue"`
}

// ExternalErrorValue represents the response in the extValeu tag.
type ExternalErrorValue struct {
	Value  interface{} `xml:"value"`
	Reason string      `xml:"reason"`
}
