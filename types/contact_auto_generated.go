package types

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in types/contact.go. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

// ContactCheckTypeIn represents a namespace agnostic version of ContactCheckType
type ContactCheckTypeIn struct {
	Check ContactCheck `xml:"command>check>check"`
}

// ContactCreateTypeIn represents a namespace agnostic version of ContactCreateType
type ContactCreateTypeIn struct {
	Create ContactCreate `xml:"command>create>create"`
}

// ContactDeleteTypeIn represents a namespace agnostic version of ContactDeleteType
type ContactDeleteTypeIn struct {
	Delete ContactDelete `xml:"command>delete>delete"`
}

// ContactInfoTypeIn represents a namespace agnostic version of ContactInfoType
type ContactInfoTypeIn struct {
	Info ContactInfo `xml:"command>info>info"`
}

// ContactTransferTypeIn represents a namespace agnostic version of ContactTransferType
type ContactTransferTypeIn struct {
	Transfer ContactTransfer `xml:"command>transfer>transfer"`
}

// ContactUpdateTypeIn represents a namespace agnostic version of ContactUpdateType
type ContactUpdateTypeIn struct {
	Update ContactUpdate `xml:"command>transfer>transfer"`
}

// ContactCheckDataTypeIn represents a namespace agnostic version of ContactCheckDataType
type ContactCheckDataTypeIn struct {
	CheckData ContactCheckData `xml:"chkData"`
}

// ContactCreateDataTypeIn represents a namespace agnostic version of ContactCreateDataType
type ContactCreateDataTypeIn struct {
	CreateData ContactCreateData `xml:"creData"`
}

// ContactInfoDataTypeIn represents a namespace agnostic version of ContactInfoDataType
type ContactInfoDataTypeIn struct {
	InfoData ContactInfoData `xml:"infData"`
}

// ContactPendingActivationNotificationDataTypeIn represents a namespace agnostic version of ContactPendingActivationNotificationDataType
type ContactPendingActivationNotificationDataTypeIn struct {
	PendingActivationNotificationData ContactPendingActivationNotificationData `xml:"panData"`
}

// ContactTransferDataTypeIn represents a namespace agnostic version of ContactTransferDataType
type ContactTransferDataTypeIn struct {
	TransferData ContactTransferData `xml:"trnData"`
}
