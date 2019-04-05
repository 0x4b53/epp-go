package types

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in types/iis.go. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

// IISExtensionCreateTypeIn represents a namespace agnostic version of IISExtensionCreateType
type IISExtensionCreateTypeIn struct {
	Create IISExtensionCreate `xml:"command>extension>create"`
}

// IISExtensionUpdateTypeIn represents a namespace agnostic version of IISExtensionUpdateType
type IISExtensionUpdateTypeIn struct {
	Update IISExtensionUpdate `xml:"command>extension>update"`
}

// IISExtensionTransferTypeIn represents a namespace agnostic version of IISExtensionTransferType
type IISExtensionTransferTypeIn struct {
	Update IISExtensionUpdate `xml:"command>extension>transfer"`
}

// IISExtensionInfoDataTypeIn represents a namespace agnostic version of IISExtensionInfoDataType
type IISExtensionInfoDataTypeIn struct {
	InfoData IISExtensionInfoData `xml:"infData"`
}
