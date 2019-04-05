package types

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in types/domain.go. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

// DomainCheckTypeIn represents a namespace agnostic version of DomainCheckType
type DomainCheckTypeIn struct {
	Check DomainCheck `xml:"command>check>check"`
}

// DomainCreateTypeIn represents a namespace agnostic version of DomainCreateType
type DomainCreateTypeIn struct {
	Create DomainCreate `xml:"command>create>create"`
}

// DomainDeleteTypeIn represents a namespace agnostic version of DomainDeleteType
type DomainDeleteTypeIn struct {
	Delete DomainDelete `xml:"command>create>delete"`
}

// DomainInfoTypeIn represents a namespace agnostic version of DomainInfoType
type DomainInfoTypeIn struct {
	Info DomainInfo `xml:"command>info>info"`
}

// DomainRenewTypeIn represents a namespace agnostic version of DomainRenewType
type DomainRenewTypeIn struct {
	Renew DomainRenew `xml:"command>renew>renew"`
}

// DomainTransferTypeIn represents a namespace agnostic version of DomainTransferType
type DomainTransferTypeIn struct {
	Transfer DomainTransfer `xml:"command>transfer>transfer"`
}

// DomainUpdateTypeIn represents a namespace agnostic version of DomainUpdateType
type DomainUpdateTypeIn struct {
	Update DomainUpdate `xml:"command>update>update"`
}

// DomainChekDataTypeIn represents a namespace agnostic version of DomainChekDataType
type DomainChekDataTypeIn struct {
	CheckData DomainCheckData `xml:"chkData"`
}

// DomainCreateDataTypeIn represents a namespace agnostic version of DomainCreateDataType
type DomainCreateDataTypeIn struct {
	CreateData DomainCreateData `xml:"creData"`
}

// DomainInfoDataTypeIn represents a namespace agnostic version of DomainInfoDataType
type DomainInfoDataTypeIn struct {
	InfoData DomainInfoData `xml:"infData"`
}

// DomainPendingActivationNotificationDataTypeIn represents a namespace agnostic version of DomainPendingActivationNotificationDataType
type DomainPendingActivationNotificationDataTypeIn struct {
	PendingActivationNotificationData DomainPendingActivationNotificationData `xml:"panData"`
}

// DomainRenewDataTypeIn represents a namespace agnostic version of DomainRenewDataType
type DomainRenewDataTypeIn struct {
	RenewData DomainRenewData `xml:"renData"`
}

// DomainTransferDataTypeIn represents a namespace agnostic version of DomainTransferDataType
type DomainTransferDataTypeIn struct {
	TransferData DomainTransferData `xml:"trnData"`
}
