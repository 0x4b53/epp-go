package types

// Name space constant for the extension.
const (
	NameSpaceDNSSEC10 = "urn:ietf:params:xml:ns:secDNS-1.0"
	NameSpaceDNSSEC11 = "urn:ietf:params:xml:ns:secDNS-1.1"
)

// DNSSECExtensionCreateType implements extension for create from secDNS-1.1
type DNSSECExtensionCreateType struct {
	Create DNSSECOrKeyData `xml:"urn:ietf:params:xml:ns:secDNS-1.1 command>extension>create"`
}

// DNSSECExtensionUpdateType implements extension for Update from secDNS-1.1
type DNSSECExtensionUpdateType struct {
	Update DNSSECExtensionUpdate `xml:"urn:ietf:params:xml:ns:secDNS-1.1 command>extension>update"`
}

// DNSSECExtensionInfoDataType represents extension for info data from secDNS-1.1
type DNSSECExtensionInfoDataType struct {
	InfoData DNSSECOrKeyData `xml:"urn:ietf:params:xml:ns:secDNS-1.1 infData"`
}

// DNSSECOrKeyData represents DNSSEC data or key data.
type DNSSECOrKeyData struct {
	MaxSignatureLife int             `xml:"maxSigLife,omitempty"`
	DNSSECData       []DNSSEC        `xml:"dsData"`
	KeyData          []DNSSECKeyData `xml:"keyData"`
}

// DNSSECExtensionUpdate implements extension for update from secDNS-1.1
type DNSSECExtensionUpdate struct {
	Remove                 DNSSECRemove    `xml:"rem,omitempty"`
	Add                    DNSSECOrKeyData `xml:"add,omitempty"`
	ChangeMaxSignatureLife int             `xml:"chg>maxSigLife,omitempty"`
	Urgent                 bool            `xml:"urgent,attr"`
}

// DNSSECRemove represents remove block for DNSSEC extension.
type DNSSECRemove struct {
	All        bool            `xml:"all,omitempty"`
	DNSSECdata []DNSSEC        `xml:"dsData"`
	KeyData    []DNSSECKeyData `xml:"keyData"`
}

// DNSSEC represents DNSSEC data.
type DNSSEC struct {
	KeyTag     uint           `xml:"keyTag"`
	Algorithm  uint           `xml:"alg"`
	DigestType uint           `xml:"digestType"`
	Digest     string         `xml:"digest"`
	KeyData    *DNSSECKeyData `xml:"keyData,omitempty"`
}

// DNSSECKeyData represents key data for DNSSEC.
type DNSSECKeyData struct {
	Flags     uint   `xml:"flags"`
	Protocol  uint   `xml:"protocol"`
	Algorithm uint   `xml:"alg"`
	PublicKey string `xml:"pubKey"`
}
