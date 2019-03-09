package types

// DomainCheck represents a check for domain(s).
type DomainCheck struct {
	Names []string `xml:"command>check>check>name,omitempty"`
}

// DomainCreate represents a domain create command.
type DomainCreate struct {
	Name        string       `xml:"command>create>create>name,omitempty"`
	Period      Period       `xml:"command>create>create>period,omitempty"`
	NameServers []NameServer `xml:"command>create>create>ns,omitempty"`
	Registrant  string       `xml:"command>create>create>registrant,omitempty"`
	Contacts    []Contact    `xml:"command>create>create>contact,omitempty"`
	AuthInfo    AuthInfo     `xml:"command>create>create>authInfo,omitempty"`
}

// DomainDelete represents a domain delete command.
type DomainDelete struct{}

// DomainInfo represents a domain info command.
type DomainInfo struct{}

// DomainRenew represents a domain renew command.
type DomainRenew struct{}

// DomainTransfer represents a domain transfer command.
type DomainTransfer struct{}

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
	HostName    string `xml:"hostName"`
	HostAddress string `xml:"hostAddr"`
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
