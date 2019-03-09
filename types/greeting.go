package types

/*
This package defines all the types used from the RFCs used to implement EPP.
Types are based of the XSDs based on the RFC but takes no
considerationDictionary to validity or constraints. This package should be used
to marshal or unmarshal EPP messages and not to veriy anything.

The package may or may not add interfaces which could then be used to implement
verification methods.
*/

import "time"

// DCPAccessType represents available DCP access types.
type DCPAccessType string

// Constants representing the string value of DCP access types.
const (
	DCPAll              DCPAccessType = "all"
	DCPNone             DCPAccessType = "none"
	DCPNull             DCPAccessType = "null"
	DCPOther            DCPAccessType = "other"
	DCPPersonal         DCPAccessType = "personal"
	DCPPersonalAndOther DCPAccessType = "personalAndOther"
)

// EPPGreeting is the type to represent a greeting from the server.
type EPPGreeting struct {
	Greeting Greeting `xml:"greeting"`
}

// Greeting represents the elements in a greeting message.
type Greeting struct {
	ServerID    string      `xml:"svID"`
	ServerDate  time.Time   `xml:"svDate"`
	ServiceMenu ServiceMenu `xml:"svcMenu"`
	DCP         DCP         `xml:"dcp"`
}

// ServiceMenu represents tags that may occur in the greeting service tag.
type ServiceMenu struct {
	Version           []string           `xml:"version"`
	Language          []string           `xml:"language"`
	ObjectURI         []string           `xml:"objURI"`
	ServiceExtentions []ServiceExtension `xml:"svcExtention"`
}

// ServiceExtension represent extension to the service.
type ServiceExtension struct {
	ExtensionURI string `xml:"extURI"`
}

// DCP (data collection policy) represents the policy declared in the greeting
// message.
type DCP struct {
	Access    DCPAccessType `xml:"access"`
	Statement DCPStatement  `xml:"statement"`
	Expiry    DCPExpiry     `xml:"expiry"`
}

// DCPExpiry represent DCP expiry.
type DCPExpiry struct {
	Absolute *time.Time `xml:"absoulte"`
	Relative string     `xml:"relative"` // Format "PnYnMnDTnHnMnS"
}

// DCPStatement represent DCP statements.
type DCPStatement struct {
	Purpose   []DCPPurpose   `xml:"purpose"`
	Recipient []DCPRecipient `xml:"recipient"`
	Retention []DCPRetention `xml:"retention"`
}

// DCPPurpose represents a DCP purposes.
type DCPPurpose struct {
	Admin   string `xml:"admin"`
	Contact string `xml:"contact"`
	Other   string `xml:"other"`
	Prov    string `xml:"prov"`
}

// DCPRecipient represents a DCP recipient.
type DCPRecipient struct {
	Other     string    `xml:"other"`
	Ours      []DCPOurs `xml:"ours"`
	Public    string    `xml:"public"`
	Same      string    `xml:"same"`
	Unrelated string    `xml:"unrelated"`
}

// DCPOurs represents the description for DCP ours.
type DCPOurs struct {
	RecipientDescription string `xml:"recDesc"`
}

// DCPRetention represents a DCP retention.
type DCPRetention struct {
	Business   string `xml:"business"`
	Indefinite string `xml:"indefinite"`
	Legal      string `xml:"legal"`
	None       string `xml:"none"`
	Stated     string `xml:"stated"`
}
