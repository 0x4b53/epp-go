package types

import "time"

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"command>check>domain:check>name"`
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name       string     `xml:"command>create>domain:create>name"`
	Period     Period     `xml:"command>create>domain:create>period,omitempty"`
	NameServer NameServer `xml:"command>create>domain:create>ns,omitempty"`
	Registrant string     `xml:"command>create>domain:create>registrant,omitempty"`
	Contacts   []Contact  `xml:"command>create>domain:create>contact,omitempty"`
	AuthInfo   AuthInfo   `xml:"command>create>domain:create>authInfo,omitempty"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct {
	Name string `xml:"command>delete>domain:delete>name"`
}

// DomainInfo represents a domain info command.
type DomainInfo struct {
	Name     InfoName `xml:"command>info>domain:info>name"`
	AuthInfo AuthInfo `xml:"command>create>domain:create>authInfo,omitempty"`
}

// InfoName ...
type InfoName struct {
	Name  string `xml:",chardata"`
	Hosts string `xml:"hosts,attr"` // TODO: Custom enum type (all, del, none, sub)
}

// DomainRenew represents a domain renew command.
type DomainRenew struct {
	Name       string    `xml:"command>renew>domain:renew>name"`
	ExpiryDate time.Time `xml:"command>renew>domain:renew>curExpDate"`
	Period     Period    `xml:"command>renew>domain:renew>period,omitempty"`
}

// DomainTransfer represents a domain transfer command.
type DomainTransfer struct {
	Name     string   `xml:"command>transfer>domain:transfer>name"`
	Period   Period   `xml:"command>transfer>domain:transfer>period,omitempty"`
	Authinfo AuthInfo `xml:"command>transfer>domain:transfer>authInfo,omitempty"`
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
	HostObject    string        `xml:"hostObj,omitempty"`
	HostAttribute HostAttribute `xml:"hostAttr,omitempty"`
}

// HostAttribute represents attributes for a host for a domain.
type HostAttribute struct {
	HostName    string   `xml:"hostName"`
	HostAddress []string `xml:"hostAddr"`
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
