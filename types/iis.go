package types

import "time"

// Name space constant for the extension.
const (
	NameSpaceIIS12 = "urn:se:iis:xml:epp:iis-1.2"
)

// IISExtensionCreateType represents the create tag from the iis-1.2 extension.
type IISExtensionCreateType struct {
	Create IISExtensionCreate `xml:"urn:se:iis:xml:epp:iis-1.2 command>extension>create"`
}

// IISExtensionUpdateType represents the update thag from iis-1.2 extension.
type IISExtensionUpdateType struct {
	Update IISExtensionUpdate `xml:"urn:se:iis:xml:epp:iis-1.2 command>extension>update"`
}

// IISExtensionTransferType represents the transfer tag from iis-1.2 extension.
type IISExtensionTransferType struct {
	Update IISExtensionUpdate `xml:"urn:se:iis:xml:epp:iis-1.2 command>extension>transfer"`
}

// IISExtensionInfoDataType represents the infData tag from iis-1.2 extension.
type IISExtensionInfoDataType struct {
	InfoData IISExtensionInfoData `xml:"urn:se:iis:xml:epp:iis-1.2 infData"`
}

// IISExtensionCreate represents the extension data for create.
type IISExtensionCreate struct {
	OrganizationNumber string `xml:"orgno,omitempty"`
	VatNumber          string `xml:"vatno,omitempty"`
}

// IISExtensionUpdate represents the extension data for update.
type IISExtensionUpdate struct {
	VatNumber    string `xml:"vatno,omitempty"`
	ClientDelete bool   `xml:"clientDelete,omitempty"`
}

// IISExtensionTransfer represents the extension data for transfer.
type IISExtensionTransfer struct {
	NameServer NameServer `xml:"ns"`
}

// IISExtensionInfoData represents the extension data for infData.
type IISExtensionInfoData struct {
	OrganizationNumber string     `xml:"orgno,omitempty"`
	VatNumber          string     `xml:"vatno,omitempty"`
	DeactivationDate   *time.Time `xml:"deactDate,omitempty"`
	DeleteDate         *time.Time `xml:"delDate,omitempty"`
	ReleaseDate        *time.Time `xml:"relDate,omitempty"`
	State              string     `xml:"state,omitempty"`
	ClientDelete       bool       `xml:"clientDelete,omitempty"`
}
