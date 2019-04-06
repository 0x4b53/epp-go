package types

import "time"

// PostalInfoType represents the typoe of a postal info.
type PostalInfoType string

// Contants represeting postal info types.
const (
	PostalInfoLocal         PostalInfoType = "loc"
	PostalInfoInternational PostalInfoType = "int"
)

// ContactStatusType represents contact status types.
type ContactStatusType string

// Constants representing contact status types.
const (
	ContactStatusClientDeleteProhibited   ContactStatusType = "clientDeleteProhibited"
	ContactStatusClientTransferProhibited ContactStatusType = "clientTransferProhibited"
	ContactStatusClientUpdateProhibited   ContactStatusType = "clientUpdateProhibited"
	ContactStatusLinked                   ContactStatusType = "linked"
	ContactStatusOk                       ContactStatusType = "ok"
	ContactStatusPendingCreate            ContactStatusType = "pendingCreate"
	ContactStatusPendingDelete            ContactStatusType = "pendingDelete"
	ContactStatusPendingTransfer          ContactStatusType = "pendingTransfer"
	ContactStatusPendingUpdate            ContactStatusType = "pendingUpdate"
	ContactStatusServerDeleteProhibited   ContactStatusType = "serverDeleteProhibited"
	ContactStatusServerTransferProhibited ContactStatusType = "serverTransferProhibited"
	ContactStatusServerUpdateProhibited   ContactStatusType = "serverUpdateProhibited"
)

// ContactTransferStatusType represents available transaction statuses.
type ContactTransferStatusType string

// Constants representing the string value of transaction status types.
const (
	ContactTransferClientApproved  ContactTransferStatusType = "clientApproved"
	ContactTransferClientCancelled ContactTransferStatusType = "clientCancelled"
	ContactTransferClientRejected  ContactTransferStatusType = "clientRejected"
	ContactTransferPending         ContactTransferStatusType = "pending"
	ContactTransferServerApproved  ContactTransferStatusType = "serverApproved"
	ContactTransferServerCancelled ContactTransferStatusType = "serverCancelled"
)

// ContactCheckType represents a contact check command.
type ContactCheckType struct {
	Check ContactCheck `xml:"urn:ietf:params:xml:ns:contact-1.0 command>check>check"`
}

// ContactCreateType represents a contact create command.
type ContactCreateType struct {
	Create ContactCreate `xml:"urn:ietf:params:xml:ns:contact-1.0 command>create>create"`
}

// ContactDeleteType represents a contact delete command.
type ContactDeleteType struct {
	Delete ContactDelete `xml:"urn:ietf:params:xml:ns:contact-1.0 command>delete>delete"`
}

// ContactInfoType represents a contact info command.
type ContactInfoType struct {
	Info ContactInfo `xml:"urn:ietf:params:xml:ns:contact-1.0 command>info>info"`
}

// ContactTransferType represents a contact transfer command.
type ContactTransferType struct {
	Transfer ContactTransfer `xml:"urn:ietf:params:xml:ns:contact-1.0 command>transfer>transfer"`
}

// ContactUpdateType represents a contact update command.
type ContactUpdateType struct {
	Update ContactUpdate `xml:"urn:ietf:params:xml:ns:contact-1.0 command>transfer>transfer"`
}

// ContactCheckDataType represents contact check data.
type ContactCheckDataType struct {
	CheckData ContactCheckData `xml:"urn:ietf:params:xml:ns:contact-1.0 chkData"`
}

// ContactCreateDataType represents contact create data.
type ContactCreateDataType struct {
	CreateData ContactCreateData `xml:"urn:ietf:params:xml:ns:contact-1.0 creData"`
}

// ContactInfoDataType represents  contact info data.
type ContactInfoDataType struct {
	InfoData ContactInfoData `xml:"urn:ietf:params:xml:ns:contact-1.0 infData"`
}

// ContactPendingActivationNotificationDataType represents contact pending
// activation notifiacation data.
type ContactPendingActivationNotificationDataType struct {
	PendingActivationNotificationData ContactPendingActivationNotificationData `xml:"urn:ietf:params:xml:ns:contact-1.0 panData"`
}

// ContactTransferDataType represents contact transfer data.
type ContactTransferDataType struct {
	TransferData ContactTransferData `xml:"urn:ietf:params:xml:ns:contact-1.0 trnData"`
}

// ContactCheck represents a check for contact(s).
type ContactCheck struct {
	Names []string `xml:"id"`
}

// ContactCreate represents a contact create command.
type ContactCreate struct {
	ID         string       `xml:"id"`
	PostalInfo []PostalInfo `xml:"postalInfo"`
	Voice      E164Type     `xml:"voice,omitempty"`
	Fax        E164Type     `xml:"fax,omitempty"`
	Email      string       `xml:"email"`
	AuthInfo   AuthInfo     `xml:"authInfo"`
	Disclose   Disclose     `xml:"disclose,omitempty"`
}

// ContactDelete represents a contact delete command.
type ContactDelete struct {
	Name string `xml:"id"`
}

// ContactInfo represents a contact info command.
type ContactInfo struct {
	Name     string   `xml:"id"`
	AuthInfo AuthInfo `xml:"authInfo,omitempty"`
}

// ContactTransfer represents a contact transfer command.
type ContactTransfer struct {
	Name     string   `xml:"id"`
	AuthInfo AuthInfo `xml:"authInfo,omitempty"`
}

// ContactUpdate represents a contact update command.
type ContactUpdate struct {
	Name   string            `xml:"id"`
	Add    *ContactAddRemove `xml:"add,omitempty"`
	Remove *ContactAddRemove `xml:"rem,omitempty"`
	Change *ContactChange    `xml:"chg>name,omitempty"`
}

// ContactCheckData represents the data returned from a contact check command.
type ContactCheckData struct {
	Name []CheckContact `xml:"cd"`
}

// ContactCreateData represents the data returned from a contact create command.
type ContactCreateData struct {
	Name       string    `xml:"id"`
	CreateDate time.Time `xml:"crDate"`
}

// ContactInfoData represents the data returned from a contact info command.
type ContactInfoData struct {
	Name         string          `xml:"id"`
	ROID         string          `xml:"roid"`
	Status       []ContactStatus `xml:"status"`
	PostalInfo   []PostalInfo    `xml:"postalInfo"`
	Voice        E164Type        `xml:"voice,omitempty"`
	Fax          E164Type        `xml:"fax,omitempty"`
	Email        string          `xml:"email"`
	ClientID     string          `xml:"clID"`
	CreateID     string          `xml:"crID"`
	CreateDate   time.Time       `xml:"crDate"`
	UpdateID     string          `xml:"upID,omitempty"`
	UpdateDate   time.Time       `xml:"upDate,omitempty"`
	TransferDate time.Time       `xml:"trDate,omitempty"`
	AuthInfo     AuthInfo        `xml:"authInfo,omitempty"`
	Disclose     Disclose        `xml:"disclose,omitempty"`
}

// ContactPendingActivationNotificationData represents the data returned from a
// contact pending activation notification command.
type ContactPendingActivationNotificationData struct {
	Name          PendingActivationNotificationName `xml:"id"`
	TransactionID string                            `xml:"paTRID"`
	Date          time.Time                         `xml:"paDate"`
}

// ContactTransferData represents the data returned from a contact transfer
// cmomand.
type ContactTransferData struct {
	Name           string                    `xml:"id"`
	TransferStatus ContactTransferStatusType `xml:"trStatus"`
	RequestingID   string                    `xml:"reID"`
	RequestingDate time.Time                 `xml:"reDate"`
	ActingID       string                    `xml:"acID"`
	ActingDate     time.Time                 `xml:"acDate"`
}

// CheckContact represents the data from a contact check command name.
type CheckContact struct {
	Name   CheckName `xml:"id"`
	Reason string    `xml:"reason,omitempty"`
}

// ContactAddRemove represents the fields that holds data to add or remove for a
// contact.
type ContactAddRemove struct {
	Status []ContactStatus `xml:"status"`
}

// ContactChange represents the data that may be changed while updating a
// contact.
type ContactChange struct {
	PostalInfo []PostalInfo `xml:"postalInfo,omitempty"`
	Voice      E164Type     `xml:"voice,omitempty"`
	Fax        E164Type     `xml:"fax,omitempty"`
	Email      string       `xml:"email,omitempty"`
	AuthInfo   AuthInfo     `xml:"authInfo,omitempty"`
	Disclose   Disclose     `xml:"disclose,omitempty,omitempty"`
}

// ContactStatus represents statuses for a contact.
type ContactStatus struct {
	Status            string            `xml:",chardata"`
	ContactStatusType ContactStatusType `xml:"s,attr"`
	Language          string            `xml:"lang,attr"`
}

// PostalInfo represents potal information for a contact.
type PostalInfo struct {
	Name         string         `xml:"name"`
	Organization string         `xml:"org,omitempty"`
	Address      Address        `xml:"addr"`
	Type         PostalInfoType `xml:"type,attr"`
}

// Address represents a full address for postal information.
type Address struct {
	Street        []string `xml:"street,omitempty"`
	City          string   `xml:"city"`
	StateProvince string   `xml:"sp,omitempty"`
	PostalCode    string   `xml:"pc,omitempty"`
	CountryCode   string   `xml:"cc"`
}

// E164Type represents the E.164 numeric plan.
type E164Type struct {
	Value string `xml:",chardata"`
	X     string `xml:"x,attr"`
}

// Disclose represents fields that may be disclosed to the public.
type Disclose struct {
	Name         InternationalOrLocalType `xml:"name,omitempty"`
	Organization InternationalOrLocalType `xml:"org,omitempty"`
	Address      InternationalOrLocalType `xml:"addr,omitempty"`
	Voice        bool                     `xml:"voice,omitempty"`
	Fax          bool                     `xml:"fax,omitempty"`
	Email        bool                     `xml:"email,omitempty"`
	Flag         bool                     `xml:"flag,attr"`
}

// InternationalOrLocalType represents a value with a type set to an available
// postal info type.
type InternationalOrLocalType struct {
	Value string         `xml:",chardata"`
	Type  PostalInfoType `xml:"type,attr"`
}
