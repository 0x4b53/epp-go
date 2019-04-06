package types

import "time"

// DomainStatusType represents available status values.
type DomainStatusType string

// Constants representing the string value of status value types.
const (
	DomainStatusClientDeleteProhibited   DomainStatusType = "clientDeleteProhibited"
	DomainStatusClientHold               DomainStatusType = "clientHold"
	DomainStatusClientRenewProhibited    DomainStatusType = "clientRenewProhibited"
	DomainStatusClientTransferProhibited DomainStatusType = "clientTransferProhibited"
	DomainStatusClientUpdateProhibited   DomainStatusType = "clientUpdateProhibited"
	DomainStatusInactive                 DomainStatusType = "inactive"
	DomainStatusOk                       DomainStatusType = "ok"
	DomainStatusPendingCreate            DomainStatusType = "pendingCreate"
	DomainStatusPendingDelete            DomainStatusType = "pendingDelete"
	DomainStatusPendingRenew             DomainStatusType = "pendingRenew"
	DomainStatusPendingTransfer          DomainStatusType = "pendingTransfer"
	DomainStatusPendingUpdate            DomainStatusType = "pendingUpdate"
	DomainStatusServerDeleteProhibited   DomainStatusType = "serverDeleteProhibited"
	DomainStatusServerHold               DomainStatusType = "serverHold"
	DomainStatusServerRenewProhibited    DomainStatusType = "serverRenewProhibited"
	DomainStatusServerTransferProhibited DomainStatusType = "serverTransferProhibited"
	DomainStatusServerUpdateProhibited   DomainStatusType = "serverUpdateProhibited"
)

// DomainHostsType represents the string value of host types.
type DomainHostsType string

// Constants representing the string value of host value types.
const (
	DomainHostsAll  DomainHostsType = "all"
	DomainHostsDel  DomainHostsType = "del"
	DomainHostsNone DomainHostsType = "none"
	DomainHostsSub  DomainHostsType = "sub"
)

// DomainTransferStatusType represents available transaction statuses.
type DomainTransferStatusType string

// Constants representing the string value of transaction status types.
const (
	DomainTransferClientApproved  DomainTransferStatusType = "clientApproved"
	DomainTransferClientCancelled DomainTransferStatusType = "clientCancelled"
	DomainTransferClientRejected  DomainTransferStatusType = "clientRejected"
	DomainTransferPending         DomainTransferStatusType = "pending"
	DomainTransferServerApproved  DomainTransferStatusType = "serverApproved"
	DomainTransferServerCancelled DomainTransferStatusType = "serverCancelled"
)

// DomainCheckType implements extension for check from domain-1.0.
type DomainCheckType struct {
	Check DomainCheck `xml:"urn:ietf:params:xml:ns:domain-1.0 command>check>check"`
}

// DomainCreateType implements extension for create from domain-1.0.
type DomainCreateType struct {
	Create DomainCreate `xml:"urn:ietf:params:xml:ns:domain-1.0 command>create>create"`
}

// DomainDeleteType implements extension for delete from domain-1.0.
type DomainDeleteType struct {
	Delete DomainDelete `xml:"urn:ietf:params:xml:ns:domain-1.0 command>create>delete"`
}

// DomainInfoType implements extension for info from domain-1.0.
type DomainInfoType struct {
	Info DomainInfo `xml:"urn:ietf:params:xml:ns:domain-1.0 command>info>info"`
}

// DomainRenewType implements extension for renew from domain-1.0.
type DomainRenewType struct {
	Renew DomainRenew `xml:"urn:ietf:params:xml:ns:domain-1.0 command>renew>renew"`
}

// DomainTransferType implements extension for transfer from domain-1.0.
type DomainTransferType struct {
	Transfer DomainTransfer `xml:"urn:ietf:params:xml:ns:domain-1.0 command>transfer>transfer"`
}

// DomainUpdateType implements extension for update from domain-1.0.
type DomainUpdateType struct {
	Update DomainUpdate `xml:"urn:ietf:params:xml:ns:domain-1.0 command>update>update"`
}

// DomainChekDataType implements extension for chekData from domain-1.0.
type DomainChekDataType struct {
	CheckData DomainCheckData `xml:"urn:ietf:params:xml:ns:domain-1.0 chkData"`
}

// DomainCreateDataType implements extension for createData from domain-1.0.
type DomainCreateDataType struct {
	CreateData DomainCreateData `xml:"urn:ietf:params:xml:ns:domain-1.0 creData"`
}

// DomainInfoDataType implements extension for infoData from domain-1.0.
type DomainInfoDataType struct {
	InfoData DomainInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}

// DomainPendingActivationNotificationDataType implements extension for
// pendingActivationNotificationData from domain-1.0.
type DomainPendingActivationNotificationDataType struct {
	PendingActivationNotificationData DomainPendingActivationNotificationData `xml:"urn:ietf:params:xml:ns:domain-1.0 panData"`
}

// DomainRenewDataType implements extension for renewData from domain-1.0.
type DomainRenewDataType struct {
	RenewData DomainRenewData `xml:"urn:ietf:params:xml:ns:domain-1.0 renData"`
}

// DomainTransferDataType implements extension for transferData from domain-1.0.
type DomainTransferDataType struct {
	TransferData DomainTransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 trnData"`
}

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"name"`
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name       string     `xml:"name"`
	Period     Period     `xml:"period,omitempty"`
	NameServer NameServer `xml:"ns,omitempty"`
	Registrant string     `xml:"registrant,omitempty"`
	Contacts   []Contact  `xml:"contact,omitempty"`
	AuthInfo   *AuthInfo  `xml:"authInfo,omitempty"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct {
	Name string `xml:"name"`
}

// DomainInfo represents a domain info command.
type DomainInfo struct {
	Name     DomainInfoName `xml:"name"`
	AuthInfo *AuthInfo      `xml:"authInfo,omitempty"`
}

// DomainInfoName represents a domain name in a domain info response.
type DomainInfoName struct {
	Name  string          `xml:",chardata"`
	Hosts DomainHostsType `xml:"hosts,attr,omitempty"`
}

// DomainRenew represents a domain renew command.
type DomainRenew struct {
	Name       string    `xml:"name"`
	ExpireDate time.Time `xml:"curExpDate"`
	Period     Period    `xml:"period,omitempty"`
}

// DomainTransfer represents a domain transfer command.
type DomainTransfer struct {
	Name     string    `xml:"command>transfer>transfer>name"`
	Period   Period    `xml:"command>transfer>transfer>period,omitempty"`
	Authinfo *AuthInfo `xml:"command>transfer>transfer>authInfo,omitempty"`
}

// DomainUpdate represents a domain update command.
type DomainUpdate struct {
	Name   string           `xml:"command>update>update>name"`
	Add    *DomainAddRemove `xml:"command>update>update>add,omitempty"`
	Remove *DomainAddRemove `xml:"command>update>update>rem,omitempty"`
	Change *DomainChange    `xml:"command>update>update>chg,omitempty"`
}

// DomainAddRemove ...
type DomainAddRemove struct {
	NameServer NameServer     `xml:"ns,omitempty"`
	Contact    []Contact      `xml:"contact,omitempty"`
	Status     []DomainStatus `xml:"status,omitempty"`
}

// DomainChange ...
type DomainChange struct {
	Registrant string    `xml:"registrant,omitempty"`
	AuthInfo   *AuthInfo `xml:"authInfo,omitempty"`
}

// Period represents the period unit and value.
type Period struct {
	Value int    `xml:",chardata"`
	Unit  string `xml:"unit,attr"`
}

// NameServer represents a name server for a domain.
type NameServer struct {
	HostObject    []string        `xml:"hostObj,omitempty"`
	HostAttribute []HostAttribute `xml:"hostAttr,omitempty"`
}

// HostAttribute represents attributes for a host for a domain.
type HostAttribute struct {
	HostName    string        `xml:"hostName"`
	HostAddress []HostAddress `xml:"hostAddr"`
}

// Contact represents a contact for a domain.
type Contact struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// AuthInfo represents authentication information used when transferring domain.
type AuthInfo struct {
	Password  string `xml:"pw,omitempty"`
	Extension string `xml:"ext,omitempty"`
}

// DomainCheckData represents the response data for a domain check command.
type DomainCheckData struct {
	CheckDomain []CheckType `xml:"cd"`
}

// DomainCreateData represents the response data for a domain create command.
type DomainCreateData struct {
	Name       string    `xml:"name"`
	CreateDate time.Time `xml:"crDate"`
	ExpireDate time.Time `xml:"exDate"`
}

// DomainInfoData represents the response data for a domain info command.
type DomainInfoData struct {
	Name         string         `xml:"name"`
	ROID         string         `xml:"roid"`
	Status       []DomainStatus `xml:"status,omitempty"`
	Registrant   string         `xml:"registrant,omitempty"`
	Contact      []Contact      `xml:"contact,omitempty"`
	NameServer   *NameServer    `xml:"ns,omitempty"`
	Host         []string       `xml:"host,omitempty"`
	ClientID     string         `xml:"clID"`
	CreateID     string         `xml:"crID,omitempty"`
	CreateDate   *time.Time     `xml:"crDate,omitempty"`
	UpdateID     string         `xml:"upID,omitempty"`
	UpdateDate   *time.Time     `xml:"upDate,omitempty"`
	ExpireDate   *time.Time     `xml:"exDate,omitempty"`
	TransferDate *time.Time     `xml:"trDate,omitempty"`
	AuthInfo     *AuthInfo      `xml:"authInfo,omitempty"`
}

// DomainPendingActivationNotificationData represents the response data for a
// domain pan command.
type DomainPendingActivationNotificationData struct {
	Name          PendingActivationNotificationName `xml:"name"`
	TransactionID string                            `xml:"paTRID"`
	Date          time.Time                         `xml:"paDate"`
}

// DomainRenewData represents the response data for a domain renew command.
type DomainRenewData struct {
	Name       string    `xml:"name"`
	ExpireDate time.Time `xml:"exDate"`
}

// DomainTransferData represents the response data for a domain transfer command.
type DomainTransferData struct {
	Name           string                   `xml:"name"`
	TransferStatus DomainTransferStatusType `xml:"trStatus"`
	RequestingID   string                   `xml:"reID"`
	RequestingDate string                   `xml:"reDate"`
	ActingID       string                   `xml:"acID"`
	ActingDate     string                   `xml:"acDate"`
	ExpireDate     string                   `xml:"exDate,omitempty"`
}

// DomainStatus represents statuses for a domain.
type DomainStatus struct {
	Status           string           `xml:",chardata"`
	DomainStatusType DomainStatusType `xml:"s,attr"`
	Language         string           `xml:"lang,attr,omitempty"`
}
