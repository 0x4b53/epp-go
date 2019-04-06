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

// HostCheckType represents a host check command.
type HostCheckType struct {
	Check HostCheck `xml:"urn:ietf:params:xml:ns:host-1.0 command>check>check"`
}

// HostCreateType represents a host create command.
type HostCreateType struct {
	Create HostCreate `xml:"urn:ietf:params:xml:ns:host-1.0 command>create>create"`
}

// HostDeleteType represents a host delete command.
type HostDeleteType struct {
	Delete HostDelete `xml:"urn:ietf:params:xml:ns:host-1.0 command>delete>delete"`
}

// HostInfoType represents a host info command.
type HostInfoType struct {
	Info HostInfo `xml:"urn:ietf:params:xml:ns:host-1.0 command>info>info"`
}

// HostUpdateType represents a host update command.
type HostUpdateType struct {
	Update HostUpdate `xml:"urn:ietf:params:xml:ns:host-1.0 command>update>update"`
}

// HostCheckDataType represents host check data.
type HostCheckDataType struct {
	CheckData HostCheckData `xml:"urn:ietf:params:xml:ns:host-1.0 chkData"`
}

// HostCreateDataType represents host create data.
type HostCreateDataType struct {
	CreateData HostCreateData `xml:"urn:ietf:params:xml:ns:host-1.0 creData"`
}

// HostInfoDataType represents host info data.
type HostInfoDataType struct {
	InfoData HostInfoData `xml:"urn:ietf:params:xml:ns:host-1.0 infData"`
}

// HostCheck represents a host check request to the EPP server.
type HostCheck struct {
	Names []string `xml:"name"`
}

// HostCreate represents a host create request to the EPP server.
type HostCreate struct {
	Name    string      `xml:"name"`
	Address HostAddress `xml:"addr,omitempty"`
}

// HostDelete represents a host delete request to the EPP server.
type HostDelete struct {
	Name string `xml:"name"`
}

// HostInfo represents a host info request to the EPP server.
type HostInfo struct {
	Name string `xml:"name"`
}

// HostUpdate represents a host update request to the EPP server.
type HostUpdate struct {
	Name   string         `xml:"name"`
	Add    *HostAddRemove `xml:"add,omitempty"`
	Remove *HostAddRemove `xml:"rem,omitempty"`
	Change string         `xml:"chg>name,omitempty"`
}

// HostCheckData represents the response for a host check command.
type HostCheckData struct {
	Name []CheckType `xml:"cd"`
}

// HostCreateData represents the response for a host create command.
type HostCreateData struct {
	Name       string    `xml:"name"`
	CreateDate time.Time `xml:"crDate"`
}

// HostInfoData represents the response for a host info command.
type HostInfoData struct {
	Name         string        `xml:"name"`
	ROID         string        `xml:"roid"`
	Status       []HostStatus  `xml:"status"`
	Address      []HostAddress `xml:"addr,omitempty"`
	ClientID     string        `xml:"clID"`
	CreateID     string        `xml:"crID"`
	CreateDate   time.Time     `xml:"crDate"`
	UpdateID     string        `xml:"upID,omitempty"`
	UpdateDate   time.Time     `xml:"upDate,omitempty"`
	TransferDate time.Time     `xml:"trDate,omitempty"`
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

// HostStatus represents statuses for a host.
type HostStatus struct {
	Status         string         `xml:",chardata"`
	HostStatusType HostStatusType `xml:"s,attr"`
	Language       string         `xml:"lang,attr"`
}
