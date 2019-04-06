package types

/*
This package defines all the types used from the RFCs used to implement EPP.
Types are based of the XSDs based on the RFC but takes no
considerationDictionary to validity or constraints. This package should be used
to marshal or unmarshal EPP messages and not to veriy anything.

The package may or may not add interfaces which could then be used to implement
verification methods.
*/

// Name space constants for the default name spaces
const (
	NameSpaceContact = "urn:ietf:params:xml:ns:contact-1.0"
	NameSpaceDomain  = "urn:ietf:params:xml:ns:domain-1.0"
	NameSpaceEPP10   = "urn:ietf:params:xml:ns:epp-1.0"
	NameSpaceHost    = "urn:ietf:params:xml:ns:host-1.0"
)

// AliasToNameSpace space will return the full name sapce for a name space alias.
func AliasToNameSpace(alias string) string {
	switch alias {
	case "contact":
		return NameSpaceContact
	case "domain":
		return NameSpaceDomain
	case "host":
		return NameSpaceHost
	}

	return ""
}

// Empty returns a non-nil value to use as an empty tag where the tag is defined
// with `omitempty` and would otherwise not be visible.
func Empty() *EmptyTag {
	return &EmptyTag{}
}

// EmptyTag represents a tag that can not have any value. This is used for
// instances to know where a tag was set or not by assigning the parent tag to a
// pointer to this type.
type EmptyTag struct{}

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
