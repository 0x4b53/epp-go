package types

// Response represents an EPP response.
type Response struct {
	Result        []Result `xml:"response>result"`
	TransactionID string   `xml:"response>trID"`
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
