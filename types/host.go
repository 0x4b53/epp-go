package types

import "time"

// IPType represents IP type, that is the IP version
type IPType string

// Contants representing allowed values for IP type.
const (
	HostIPv4 IPType = "v4"
	HostIPv6 IPType = "v6"
)

// HostStatusType represents available status values.
type HostStatusType string

// Constants representing the string value of status value types.
const (
	HostStatusClientDeleteProhibited HostStatusType = "clientDeleteProhibited"
	HostStatusClientUpdateProhibited HostStatusType = "clientUpdateProhibited"
	HostStatusLinked                 HostStatusType = "linked"
	HostStatusOk                     HostStatusType = "ok"
	HostStatusPendingCreate          HostStatusType = "pendingCreate"
	HostStatusPendingDelete          HostStatusType = "pendingDelete"
	HostStatusPendingTransfer        HostStatusType = "pendingTransfer"
	HostStatusPendingUpdate          HostStatusType = "pendingUpdate"
	HostStatusServerDeleteProhibited HostStatusType = "serverDeleteProhibited"
	HostStatusServerUpdateProhibited HostStatusType = "serverUpdateProhibited"
)

// HostCheck represents a host check request to the EPP server.
type HostCheck struct {
	Names []string `xml:"command>check>check>name"`
}

// HostCreate represents a host create request to the EPP server.
type HostCreate struct {
	Name    string      `xml:"command>create>create>name"`
	Address HostAddress `xml:"command>create>create>addr,omitempty"`
}

// HostDelete represents a host delete request to the EPP server.
type HostDelete struct {
	Name string `xml:"command>delete>delete>name"`
}

// HostInfo represents a host info request to the EPP server.
type HostInfo struct {
	Name string `xml:"command>info>hot:info>name"`
}

// HostUpdate represents a host update request to the EPP server.
type HostUpdate struct {
	Name   string         `xml:"command>update>update>name"`
	Add    *HostAddRemove `xml:"command>update>update>add,omitempty"`
	Remove *HostAddRemove `xml:"command>update>update>rem,omitempty"`
	Change string         `xml:"command>update>update>chg>name,omitempty"`
}

// HostAddRemove represents data that can be added or removed while updating a
// domain.
type HostAddRemove struct {
	Address []HostAddress `xml:"addr,omitempty"`
}

// HostAddress represents an IP address beloning to a host.
type HostAddress struct {
	Address string `xml:",chardata,omitempty"`
	IP      IPType `xml:"ip,attr"`
}

// HostCheckData represents the response for a host check command.
type HostCheckData struct {
	Name []CheckType `xml:"chkData>cd"`
}

// HostCreateData represents the response for a host create command.
type HostCreateData struct {
	Name       string    `xml:"creData>name"`
	CreateDate time.Time `xml:"creData>crDate"`
}

// HostInfoData represents the response for a host info command.
type HostInfoData struct {
	Name         string        `xml:"infData>name"`
	ROID         string        `xml:"infData>roid"`
	Status       []HostStatus  `xml:"infData>status"`
	Address      []HostAddress `xml:"infData>addr,omitempty"`
	ClientID     string        `xml:"infData>clID"`
	CreateID     string        `xml:"infData>crID"`
	CreateDate   time.Time     `xml:"infData>crDate"`
	UpdateID     string        `xml:"infData>upID,omitempty"`
	UpdateDate   time.Time     `xml:"infData>upDate,omitempty"`
	TransferDate time.Time     `xml:"infData>trDate,omitempty"`
}

// HostStatus represents statuses for a host.
type HostStatus struct {
	Status         string         `xml:",chardata"`
	HostStatusType HostStatusType `xml:"s,attr"`
	Language       string         `xml:"lang,attr"`
}
