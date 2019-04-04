package types

import "time"

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
	Language          []string           `xml:"lang"`
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
	Access    DCPAccess    `xml:"access"`
	Statement DCPStatement `xml:"statement"`
	Expiry    *DCPExpiry   `xml:"expiry,omitempty"`
}

// DCPAccess represents the access type.
type DCPAccess struct {
	All              *EmptyTag `xml:"all,omitempty"`
	None             *EmptyTag `xml:"none,omitempty"`
	Null             *EmptyTag `xml:"null,omitempty"`
	Other            *EmptyTag `xml:"other,omitempty"`
	Personal         *EmptyTag `xml:"personal,omitempty"`
	PersonalAndOther *EmptyTag `xml:"personalAndOther,omitempty"`
}

// DCPExpiry represent DCP expiry.
type DCPExpiry struct {
	Absolute *time.Time `xml:"absoulte"`
	Relative string     `xml:"relative,omitempty"` // Format "PnYnMnDTnHnMnS"
}

// DCPStatement represent DCP statements.
type DCPStatement struct {
	Purpose   DCPPurpose   `xml:"purpose"`
	Recipient DCPRecipient `xml:"recipient"`
	Retention DCPRetention `xml:"retention"`
}

// DCPPurpose represents a DCP purposes.
type DCPPurpose struct {
	Admin   *EmptyTag `xml:"admin,omitempty"`
	Contact *EmptyTag `xml:"contact,omitempty"`
	Other   *EmptyTag `xml:"other,omitempty"`
	Prov    *EmptyTag `xml:"prov,omitempty"`
}

// DCPRecipient represents a DCP recipient.
type DCPRecipient struct {
	Other     *EmptyTag `xml:"other"`
	Ours      []DCPOurs `xml:"ours"`
	Public    *EmptyTag `xml:"public"`
	Same      *EmptyTag `xml:"same"`
	Unrelated *EmptyTag `xml:"unrelated"`
}

// DCPOurs represents the description for DCP ours.
type DCPOurs struct {
	RecipientDescription string `xml:"recDesc,omitempty"`
}

// DCPRetention represents a DCP retention.
type DCPRetention struct {
	Business   *EmptyTag `xml:"business"`
	Indefinite *EmptyTag `xml:"indefinite"`
	Legal      *EmptyTag `xml:"legal"`
	None       *EmptyTag `xml:"none"`
	Stated     *EmptyTag `xml:"stated"`
}
