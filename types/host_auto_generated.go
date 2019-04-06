package types

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in types/host.go. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

// HostCheckTypeIn represents a namespace agnostic version of HostCheckType
type HostCheckTypeIn struct {
	Check HostCheck `xml:"command>check>check"`
}

// HostCreateTypeIn represents a namespace agnostic version of HostCreateType
type HostCreateTypeIn struct {
	Create HostCreate `xml:"command>create>create"`
}

// HostDeleteTypeIn represents a namespace agnostic version of HostDeleteType
type HostDeleteTypeIn struct {
	Delete HostDelete `xml:"command>delete>delete"`
}

// HostInfoTypeIn represents a namespace agnostic version of HostInfoType
type HostInfoTypeIn struct {
	Info HostInfo `xml:"command>info>info"`
}

// HostUpdateTypeIn represents a namespace agnostic version of HostUpdateType
type HostUpdateTypeIn struct {
	Update HostUpdate `xml:"command>update>update"`
}

// HostCheckDataTypeIn represents a namespace agnostic version of HostCheckDataType
type HostCheckDataTypeIn struct {
	CheckData HostCheckData `xml:"chkData"`
}

// HostCreateDataTypeIn represents a namespace agnostic version of HostCreateDataType
type HostCreateDataTypeIn struct {
	CreateData HostCreateData `xml:"creData"`
}

// HostInfoDataTypeIn represents a namespace agnostic version of HostInfoDataType
type HostInfoDataTypeIn struct {
	InfoData HostInfoData `xml:"infData"`
}
