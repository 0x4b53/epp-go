package types

// HostCheck ...
type HostCheck struct {
	Names []string `xml:"command>check>host:check>name"`
}

// HostCreate ...
type HostCreate struct {
	Name    HostName    `xml:"command>create>host:create>name,omitempty"`
	Address HostAddress `xml:"command>create>host:create>addr,omitempty"`
}

// HostDelete ...
type HostDelete struct {
	Name string `xml:"command>delete>host:delete>name"`
}

// HostInfo ...
type HostInfo struct {
	Name string `xml:"command>info>hot:info>name"`
}

// HostUpdate ...
type HostUpdate struct {
	Name   string        `xml:"command>update>host:update>name"`
	Add    []HostAddress `xml:"command>update>host:update>add,omitempty"`
	Remove []HostAddress `xml:"command>update>host:update>rem,omitempty"`
	Change []HostName    `xml:"command>update>host:update>chg>name,omitempty"`
}

// HostName ...
type HostName struct {
	Name     string `xml:",chardata"`
	PaResult bool   `xml:"paResult,attr,omitempty"`
}

// HostAddress ...
type HostAddress struct {
	Address string `xml:",chardata"`
	IP      string `xml:"addr,attr"`
}
