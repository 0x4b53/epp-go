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
	Names []string `xml:"command>check>host:check>name"`
}

// HostCreate represents a host create request to the EPP server.
type HostCreate struct {
	Name    string               `xml:"command>create>host:create>host:name"`
	Address HostAddRemoveAddress `xml:"command>create>host:create>host:addr,omitempty"`
}

// HostDelete represents a host delete request to the EPP server.
type HostDelete struct {
	Name string `xml:"command>delete>host:delete>name"`
}

// HostInfo represents a host info request to the EPP server.
type HostInfo struct {
	Name string `xml:"command>info>hot:info>name"`
}

// HostUpdate represents a host update request to the EPP server.
type HostUpdate struct {
	Name   string         `xml:"command>update>host:update>host:name"`
	Add    *HostAddRemove `xml:"command>update>host:update>host:add,omitempty"`
	Remove *HostAddRemove `xml:"command>update>host:update>host:rem,omitempty"`
	Change string         `xml:"command>update>host:update>host:chg>host:name,omitempty"`
}

// HostAddRemove ...
type HostAddRemove struct {
	Address []HostAddRemoveAddress `xml:"host:addr,omitempty"`
}

// HostAddRemoveAddress ...
type HostAddRemoveAddress struct {
	Address string `xml:",chardata,omitempty"`
	IP      IPType `xml:"ip,attr"`
}

// HostCheckData ...
type HostCheckData struct {
	Name []CheckHost `xml:"host:chkData>host:cd"`
}

// HostCreateData ...
type HostCreateData struct {
	Name       string    `xml:"host:creData>host:name"`
	CreateDate time.Time `xml:"host:creData>host:crDate"`
}

// HostInfoData ...
type HostInfoData struct {
	Name         string                 `xml:"host:infData>host:name"`
	ROID         string                 `xml:"host:infData>host:roid"`
	Status       []HostStatus           `xml:"host:infData>host:status"`
	Address      []HostAddRemoveAddress `xml:"host:infData>host:addr,omitempty"`
	ClientID     string                 `xml:"host:infData>host:clID"`
	CreateID     string                 `xml:"host:infData>host:crID"`
	CreateDate   time.Time              `xml:"host:infData>host:crDate"`
	UpdateID     string                 `xml:"host:infData>host:upID,omitempty"`
	UpdateDate   time.Time              `xml:"host:infData>host:upDate,omitempty"`
	TransferDate time.Time              `xml:"host:infData>host:trDate,omitempty"`
}

// HostPendingActivationNotificationData ...
type HostPendingActivationNotificationData struct {
	Name          HostPendingActivationNotificationName `xml:"host:panData>host:name"`
	TransactionID string                                `xml:"host:panData>host:paTRID"`
	Date          time.Time                             `xml:"host:panData>host:paDate"`
}

// CheckHost ...
type CheckHost struct {
	Name   CheckHostName `xml:"host:name"`
	Reason string        `xml:"host:reason,omitempty"`
}

// CheckHostName ...
type CheckHostName struct {
	Value     string `xml:",chardata"`
	Available bool   `xml:"avail,attr"`
}

// HostStatus represents statuses for a domain.
type HostStatus struct {
	Status         string         `xml:",chardata"`
	HostStatusType HostStatusType `xml:"s,attr"`
	Language       string         `xml:"lang,attr"`
}

// HostPendingActivationNotificationName represents the host name tag in a pending
// activation notification response.
type HostPendingActivationNotificationName struct {
	Name                    string `xml:",chardata"`
	PendingActivationResult bool   `xml:"paResult,attr,omitempty"`
}
