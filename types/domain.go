package types

import "time"

// StatusValueType represents available status values.
type StatusValueType string

// COnstants representing the string value of status value types.
const (
	ClientDeleteProhibited   StatusValueType = "clientDeleteProhibited"
	ClientHold               StatusValueType = "clientHold"
	ClientRenewProhibited    StatusValueType = "clientRenewProhibited"
	ClientTransferProhibited StatusValueType = "clientTransferProhibited"
	ClientUpdateProhibited   StatusValueType = "clientUpdateProhibited"
	Inactive                 StatusValueType = "inactive"
	Ok                       StatusValueType = "ok"
	PendingCreate            StatusValueType = "pendingCreate"
	PendingDelete            StatusValueType = "pendingDelete"
	PendingRenew             StatusValueType = "pendingRenew"
	PendingTransfer          StatusValueType = "pendingTransfer"
	PendingUpdate            StatusValueType = "pendingUpdate"
	ServerDeleteProhibited   StatusValueType = "serverDeleteProhibited"
	ServerHold               StatusValueType = "serverHold"
	ServerRenewProhibited    StatusValueType = "serverRenewProhibited"
	ServerTransferProhibited StatusValueType = "serverTransferProhibited"
	ServerUpdateProhibited   StatusValueType = "serverUpdateProhibited"
)

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"command>check>domain:check>domain:name"`
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name       string     `xml:"command>create>domain:create>domain:name"`
	Period     Period     `xml:"command>create>domain:create>domain:period,omitempty"`
	NameServer NameServer `xml:"command>create>domain:create>domain:ns,omitempty"`
	Registrant string     `xml:"command>create>domain:create>domain:registrant,omitempty"`
	Contacts   []Contact  `xml:"command>create>domain:create>domain:contact,omitempty"`
	AuthInfo   AuthInfo   `xml:"command>create>domain:create>domain:authInfo,omitempty"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct {
	Name string `xml:"command>delete>domain:delete>domain:name"`
}

// DomainInfo represents a domain info command.
type DomainInfo struct {
	Name     InfoName `xml:"command>info>domain:info>domain:name"`
	AuthInfo AuthInfo `xml:"command>create>domain:create>domain:authInfo,omitempty"`
}

// InfoName ...
type InfoName struct {
	Name  string `xml:",chardata"`
	Hosts string `xml:"hosts,attr"` // TODO: Custom enum type (all, del, none, sub)
}

// DomainRenew represents a domain renew command.
type DomainRenew struct {
	Name       string    `xml:"command>renew>domain:renew>domain:name"`
	ExpiryDate time.Time `xml:"command>renew>domain:renew>domain:curExpDate"`
	Period     Period    `xml:"command>renew>domain:renew>domain:period,omitempty"`
}

// DomainTransfer represents a domain transfer command.
type DomainTransfer struct {
	Name     string   `xml:"command>transfer>domain:transfer>domain:name"`
	Period   Period   `xml:"command>transfer>domain:transfer>domain:period,omitempty"`
	Authinfo AuthInfo `xml:"command>transfer>domain:transfer>domain:authInfo,omitempty"`
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
	HostObject    string        `xml:"domain:hostObj,omitempty"`
	HostAttribute HostAttribute `xml:"domain:hostAttr,omitempty"`
}

// HostAttribute represents attributes for a host for a domain.
type HostAttribute struct {
	HostName    string   `xml:"domain:hostName"`
	HostAddress []string `xml:"domain:hostAddr"`
}

// Contact represents a contact for a domain.
type Contact struct {
	Name string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// AuthInfo represents authentication information used when transferring domain.
type AuthInfo struct {
	Password  string `xml:"domain:pw,omitempty"`
	Extension string `xml:"domain:ext,omitempty"`
}

// DomainCheckData represents the response data for a domain check command.
type DomainCheckData struct {
	CheckDomain []CheckDomain `xml:"domain:chkData>domain:cd"`
}

// DomainCreateData represents the response data for a domain create command.
type DomainCreateData struct {
	Name       string    `xml:"domain:creData>domain:name"`
	CreateDate time.Time `xml:"domain:creData>domain:crDate"`
	ExpireDate time.Time `xml:"domain:creData>domain:exDate"`
}

// DomainInfoData represents the response data for a domain info command.
type DomainInfoData struct {
	Name         string     `xml:"domain:infData>domain:name"`
	ROID         string     `xml:"domain:infData>domain:roid"`
	Status       []string   `xml:"domain:infData>domain:status,omitempty"` // TODO: Enum
	Registrant   string     `xml:"domain:infData>domain:registrant,omitempty"`
	Contact      []Contact  `xml:"domain:infData>domain:contact,omitempty"`
	NameServer   NameServer `xml:"domain:infData>domain:ns"`
	Host         []string   `xml:"domain:infData>domain:host,omitempty"`
	ClientID     string     `xml:"domain:infData>domain:clID"`
	CreateID     string     `xml:"domain:infData>domain:crID,omitempty"`
	CreateDate   time.Time  `xml:"domain:infData>domain:crDate,omitempty"`
	UpdateID     time.Time  `xml:"domain:infData>domain:upID,omitempty"`
	UpdateDate   time.Time  `xml:"domain:infData>domain:upDate,omitempty"`
	ExpireDate   time.Time  `xml:"domain:infData>domain:exDate,omitempty"`
	TransferDate time.Time  `xml:"domain:infData>domain:trDate,omitempty"`
	AuthInfo     AuthInfo   `xml:"domain:infData>domain:authInfo,omitempty"`
}

// DomainPendingActivationNotificationData represents the response data for a
// domain pan command.
type DomainPendingActivationNotificationData struct {
	Name                           DomainName `xml:"domain:panData>domain:name"`
	PendingActivationTransactionID string     `xml:"domain:panData>domain:paTRID"`
	PendingActivationDate          time.Time  `xml:"domain:panData>domain:paDate"`
}

// DomainRenewData represents the response data for a domain renew command.
type DomainRenewData struct{}

// DomainTransferData represents the response data for a domain transfer command.
type DomainTransferData struct{}

// DomainName ...
type DomainName struct {
	Name                    string `xml:",chardata"`
	PendingActivationResult bool   `xml:"paResult,attr,omitempty"`
}

// CheckDomain ...
type CheckDomain struct {
	Name   CheckDomainName `xml:"domain:name"`
	Reason string          `xml:"domain:reason,omitempty"`
}

// CheckDomainName ...
type CheckDomainName struct {
	Value     string `xml:",chardata"`
	Available bool   `xml:"avail,attr"`
}
