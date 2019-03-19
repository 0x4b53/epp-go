package types

import (
	"time"
)

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

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"command>check>check>name"`
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name       string     `xml:"command>create>create>name"`
	Period     Period     `xml:"command>create>create>period,omitempty"`
	NameServer NameServer `xml:"command>create>create>ns,omitempty"`
	Registrant string     `xml:"command>create>create>registrant,omitempty"`
	Contacts   []Contact  `xml:"command>create>create>contact,omitempty"`
	AuthInfo   AuthInfo   `xml:"command>create>create>authInfo,omitempty"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct {
	Name string `xml:"command>delete>delete>name"`
}

// DomainInfo represents a domain info command.
type DomainInfo struct {
	Name     DomainInfoName `xml:"command>info>info>name"`
	AuthInfo AuthInfo       `xml:"command>create>create>authInfo,omitempty"`
}

// DomainInfoName represents a domain name in a domain info response.
type DomainInfoName struct {
	Name  string          `xml:",chardata"`
	Hosts DomainHostsType `xml:"hosts,attr"`
}

// DomainRenew represents a domain renew command.
type DomainRenew struct {
	Name       string    `xml:"command>renew>renew>name"`
	ExpireDate time.Time `xml:"command>renew>renew>curExpDate"`
	Period     Period    `xml:"command>renew>renew>period,omitempty"`
}

// DomainTransfer represents a domain transfer command.
type DomainTransfer struct {
	Name     string   `xml:"command>transfer>transfer>name"`
	Period   Period   `xml:"command>transfer>transfer>period,omitempty"`
	Authinfo AuthInfo `xml:"command>transfer>transfer>authInfo,omitempty"`
}

// DomainUpdate represents a domain update command.
type DomainUpdate struct{}

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
	CheckDomain []CheckDomain `xml:"chkData>cd"`
}

// DomainCreateData represents the response data for a domain create command.
type DomainCreateData struct {
	Name       string    `xml:"creData>name"`
	CreateDate time.Time `xml:"creData>crDate"`
	ExpireDate time.Time `xml:"creData>exDate"`
}

// DomainInfoData represents the response data for a domain info command.
type DomainInfoData struct {
	Name         string         `xml:"infData>name"`
	ROID         string         `xml:"infData>roid"`
	Status       []DomainStatus `xml:"infData>status,omitempty"`
	Registrant   string         `xml:"infData>registrant,omitempty"`
	Contact      []Contact      `xml:"infData>contact,omitempty"`
	NameServer   NameServer     `xml:"infData>ns"`
	Host         []string       `xml:"infData>host,omitempty"`
	ClientID     string         `xml:"infData>clID"`
	CreateID     string         `xml:"infData>crID,omitempty"`
	CreateDate   time.Time      `xml:"infData>crDate,omitempty"`
	UpdateID     time.Time      `xml:"infData>upID,omitempty"`
	UpdateDate   time.Time      `xml:"infData>upDate,omitempty"`
	ExpireDate   time.Time      `xml:"infData>exDate,omitempty"`
	TransferDate time.Time      `xml:"infData>trDate,omitempty"`
	AuthInfo     AuthInfo       `xml:"infData>authInfo,omitempty"`
}

// DomainPendingActivationNotificationData represents the response data for a
// domain pan command.
type DomainPendingActivationNotificationData struct {
	Name          DomainPendingActivationNotificationName `xml:"panData>name"`
	TransactionID string                                  `xml:"panData>paTRID"`
	Date          time.Time                               `xml:"panData>paDate"`
}

// DomainRenewData represents the response data for a domain renew command.
type DomainRenewData struct {
	Name       string    `xml:"renData>name"`
	ExpireDate time.Time `xml:"renData>exDate"`
}

// DomainTransferData represents the response data for a domain transfer command.
type DomainTransferData struct {
	Name           string                   `xml:"trnData>name"`
	TransferStatus DomainTransferStatusType `xml:"trnData>trStatus"`
	RequestingID   string                   `xml:"trnData>reID"`
	RequestingDate string                   `xml:"trnData>reDate"`
	ActingID       string                   `xml:"trnData>acID"`
	ActingDate     string                   `xml:"trnData>acDate"`
	ExpireDate     string                   `xml:"trnData>exDate,omitempty"`
}

// DomainPendingActivationNotificationName represents the domain name tag in a
// pending activation notificateion response.
type DomainPendingActivationNotificationName struct {
	Name                    string `xml:",chardata"`
	PendingActivationResult bool   `xml:"paResult,attr,omitempty"`
}

// CheckDomain ...
type CheckDomain struct {
	Name   CheckDomainName `xml:"name"`
	Reason string          `xml:"reason,omitempty"`
}

// CheckDomainName ...
type CheckDomainName struct {
	Value     string `xml:",chardata"`
	Available bool   `xml:"avail,attr"`
}

// DomainStatus represents statuses for a domain.
type DomainStatus struct {
	Status           string           `xml:",chardata"`
	DomainStatusType DomainStatusType `xml:"s,attr"`
	Language         string           `xml:"lang,attr"`
}
