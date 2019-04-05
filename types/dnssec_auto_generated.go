package types

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in types/dnssec.go. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

// DNSSECExtensionCreateTypeIn represents a namespace agnostic version of DNSSECExtensionCreateType
type DNSSECExtensionCreateTypeIn struct {
	Create DNSSECOrKeyData `xml:"command>extension>create"`
}

// DNSSECExtensionUpdateTypeIn represents a namespace agnostic version of DNSSECExtensionUpdateType
type DNSSECExtensionUpdateTypeIn struct {
	Update DNSSECExtensionUpdate `xml:"command>extension>update"`
}

// DNSSECExtensionInfoDataTypeIn represents a namespace agnostic version of DNSSECExtensionInfoDataType
type DNSSECExtensionInfoDataTypeIn struct {
	InfoData DNSSECOrKeyData `xml:"infData"`
}
