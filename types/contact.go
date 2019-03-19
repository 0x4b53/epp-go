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

// ContactCheck represents a check for contact(s).
type ContactCheck struct {
	Names []string `xml:"command>check>check>id"`
}

// ContactCreate represents a contact create command.
type ContactCreate struct {
	ID         string       `xml:"command>create>create>id"`
	PostalInfo []PostalInfo `xml:"command>create>create>postalInfo"`
	Voice      E164Type     `xml:"command>create>create>voice,omitempty"`
	Fax        E164Type     `xml:"command>create>create>fax,omitempty"`
	Email      string       `xml:"command>create>create>email"`
	AuthInfo   AuthInfo     `xml:"command>create>create>authInfo"`
	Disclose   Disclose     `xml:"command>create>create>disclose,omitempty"`
}

// ContactDelete represents a contact delete command.
type ContactDelete struct {
	Name string `xml:"command>delete>delete>id"`
}

// ContactInfo represents a contact info command.
type ContactInfo struct {
	Name     string   `xml:"command>info>info>id"`
	AuthInfo AuthInfo `xml:"command>info>info>authInfo,omitempty"`
}

// ContactTransfer represents a contact transfer command.
type ContactTransfer struct {
	Name     string   `xml:"command>info>info>id"`
	AuthInfo AuthInfo `xml:"command>info>info>authInfo,omitempty"`
}

// ContactUpdate represents a contact update command.
type ContactUpdate struct {
	Name   string            `xml:"command>update>update>id"`
	Add    *ContactAddRemove `xml:"command>update>update>add,omitempty"`
	Remove *ContactAddRemove `xml:"command>update>update>rem,omitempty"`
	Change *ContactChange    `xml:"command>update>update>chg>name,omitempty"`
}

// ContactCheckData represents the data returned from a contact check command.
type ContactCheckData struct {
	Name []CheckContact `xml:"chkData>cd"`
}

// ContactCreateData represents the data returned from a contact create command.
type ContactCreateData struct {
	Name       string    `xml:"creDataidname"`
	CreateDate time.Time `xml:"creData>crDate"`
}

// ContactInfoData represents the data returned from a contact info command.
type ContactInfoData struct {
	Name         string          `xml:"infData>id"`
	ROID         string          `xml:"infData>roid"`
	Status       []ContactStatus `xml:"infData>status"`
	PostalInfo   []PostalInfo    `xml:"infData>postalInfo"`
	Voice        E164Type        `xml:"infData>voice,omitempty"`
	Fax          E164Type        `xml:"infData>fax,omitempty"`
	Email        string          `xml:"infData>email"`
	ClientID     string          `xml:"infData>clID"`
	CreateID     string          `xml:"infData>crID"`
	CreateDate   time.Time       `xml:"infData>crDate"`
	UpdateID     string          `xml:"infData>upID,omitempty"`
	UpdateDate   time.Time       `xml:"infData>upDate,omitempty"`
	TransferDate time.Time       `xml:"infData>trDate,omitempty"`
	AuthInfo     AuthInfo        `xml:"infData>authInfo,omitempty"`
	Disclose     Disclose        `xml:"infData>disclose,omitempty"`
}

// ContactPendingActivationNotificationData represents the data returned from a
// contact pending activation notification command.
type ContactPendingActivationNotificationData struct {
	Name          PendingActivationNotificationName `xml:"panData>id"`
	TransactionID string                            `xml:"panData>paTRID"`
	Date          time.Time                         `xml:"panData>paDate"`
}

// ContactTransferData represents the data returned from a contact transfer
// cmomand.
type ContactTransferData struct {
	Name           string                    `xml:"trnData>id"`
	TransferStatus ContactTransferStatusType `xml:"trnData>trStatus"`
	RequestingID   string                    `xml:"trnData>reID"`
	RequestingDate time.Time                 `xml:"trnData>reDate"`
	ActingID       string                    `xml:"trnData>acID"`
	ActingDate     time.Time                 `xml:"trnData>acDate"`
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
