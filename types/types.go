package types

/*
This package defines all the types used from the RFCs used to implement EPP.
Types are based of the XSDs based on the RFC but takes no
considerationDictionary to validity or constraints. This package should be used
to marshal or unmarshal EPP messages and not to veriy anything.

The package may or may not add interfaces which could then be used to implement
verification methods.
*/

// CheckType represents the data from any kind of check command.
type CheckType struct {
	Name   CheckName `xml:"name"`
	Reason string    `xml:"reason,omitempty"`
}

// CheckName represents the name in a check command.
type CheckName struct {
	Value     string `xml:",chardata"`
	Available bool   `xml:"avail,attr"`
}

// PendingActivationNotificationName represents the name in pending activation
// notification data sets.
type PendingActivationNotificationName struct {
	Name                    string `xml:",chardata"`
	PendingActivationResult bool   `xml:"paResult,attr,omitempty"`
}
