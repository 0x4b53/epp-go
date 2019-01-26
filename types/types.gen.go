package ws

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"time"
)

type AddRemType struct {
	Addr   []AddrType   `xml:"urn:ietf:params:xml:ns:host-1.0 addr,omitempty"`
	Status []StatusType `xml:"urn:ietf:params:xml:ns:host-1.0 status,omitempty"`
}

// Must be at least 3 items long
type AddrStringType string

type AddrType struct {
	Street []string       `xml:"urn:ietf:params:xml:ns:contact-1.0 street,omitempty"`
	City   PostalLineType `xml:"urn:ietf:params:xml:ns:contact-1.0 city"`
	Sp     string         `xml:"urn:ietf:params:xml:ns:contact-1.0 sp,omitempty"`
	Pc     string         `xml:"urn:ietf:params:xml:ns:contact-1.0 pc,omitempty"`
	Cc     CcType         `xml:"urn:ietf:params:xml:ns:contact-1.0 cc"`
}

type AuthIDType struct {
	Id       ClIDType     `xml:"urn:ietf:params:xml:ns:contact-1.0 id"`
	AuthInfo AuthInfoType `xml:"urn:ietf:params:xml:ns:contact-1.0 authInfo,omitempty"`
}

type AuthInfoChgType struct {
	Pw   PwAuthInfoType  `xml:"urn:ietf:params:xml:ns:domain-1.0 pw"`
	Ext  ExtAuthInfoType `xml:"urn:ietf:params:xml:ns:domain-1.0 ext"`
	Null string          `xml:"urn:ietf:params:xml:ns:domain-1.0 null"`
}

type AuthInfoType struct {
	Pw  PwAuthInfoType  `xml:"urn:ietf:params:xml:ns:contact-1.0 pw"`
	Ext ExtAuthInfoType `xml:"urn:ietf:params:xml:ns:contact-1.0 ext"`
}

// May be no more than 2 items long
type CcType string

type CheckIDType struct {
	ClIDType ClIDType `xml:",chardata"`
	Avail    bool     `xml:"avail,attr"`
}

type CheckNameType struct {
	LabelType LabelType `xml:",chardata"`
	Avail     bool      `xml:"avail,attr"`
}

type CheckType struct {
	Name   CheckNameType `xml:"urn:ietf:params:xml:ns:host-1.0 name"`
	Reason ReasonType    `xml:"urn:ietf:params:xml:ns:host-1.0 reason,omitempty"`
}

type ChgPostalInfoType struct {
	Name PostalLineType     `xml:"urn:ietf:params:xml:ns:contact-1.0 name,omitempty"`
	Org  string             `xml:"urn:ietf:params:xml:ns:contact-1.0 org,omitempty"`
	Addr AddrType           `xml:"urn:ietf:params:xml:ns:contact-1.0 addr,omitempty"`
	Type PostalInfoEnumType `xml:"type,attr"`
}

type ChgType struct {
	MaxSigLife int `xml:"urn:ietf:params:xml:ns:secDNS-1.1 maxSigLife,omitempty"`
}

type ChkDataType struct {
	Cd []CheckType `xml:"urn:ietf:params:xml:ns:host-1.0 cd"`
}

// Must be at least 3 items long
type ClIDType string

type CommandType struct {
	Check     ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 check"`
	Create    ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 create"`
	Delete    ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 delete"`
	Info      ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 info"`
	Login     LoginType      `xml:"urn:ietf:params:xml:ns:epp-1.0 login"`
	Logout    string         `xml:"urn:ietf:params:xml:ns:epp-1.0 logout"`
	Poll      PollType       `xml:"urn:ietf:params:xml:ns:epp-1.0 poll"`
	Renew     ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 renew"`
	Transfer  TransferType   `xml:"urn:ietf:params:xml:ns:epp-1.0 transfer"`
	Update    ReadWriteType  `xml:"urn:ietf:params:xml:ns:epp-1.0 update"`
	Extension ExtAnyType     `xml:"urn:ietf:params:xml:ns:epp-1.0 extension,omitempty"`
	ClTRID    TrIDStringType `xml:"urn:ietf:params:xml:ns:epp-1.0 clTRID,omitempty"`
}

// May be one of admin, billing, tech
type ContactAttrType string

type ContactType struct {
	ClIDType ClIDType        `xml:",chardata"`
	Type     ContactAttrType `xml:"type,attr,omitempty"`
}

type CreDataType struct {
	Name   LabelType `xml:"urn:ietf:params:xml:ns:host-1.0 name"`
	CrDate time.Time `xml:"urn:ietf:params:xml:ns:host-1.0 crDate"`
}

func (t *CreDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CreDataType
	var layout struct {
		*T
		CrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:host-1.0 crDate"`
	}
	layout.T = (*T)(t)
	layout.CrDate = (*xsdDateTime)(&layout.T.CrDate)
	return e.EncodeElement(layout, start)
}
func (t *CreDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CreDataType
	var overlay struct {
		*T
		CrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:host-1.0 crDate"`
	}
	overlay.T = (*T)(t)
	overlay.CrDate = (*xsdDateTime)(&overlay.T.CrDate)
	return d.DecodeElement(&overlay, &start)
}

type CreateNotifyType struct {
	InfData InfDataType `xml:"urn:se:iis:xml:epp:iis-1.2 infData"`
}

type CreateType struct {
	Orgno OrgnoType `xml:"urn:se:iis:xml:epp:iis-1.2 orgno"`
	Vatno string    `xml:"urn:se:iis:xml:epp:iis-1.2 vatno,omitempty"`
}

type CredsOptionsType struct {
	Version VersionType `xml:"urn:ietf:params:xml:ns:epp-1.0 version"`
	Lang    string      `xml:"urn:ietf:params:xml:ns:epp-1.0 lang"`
}

type DcpAccessType struct {
	All              string `xml:"urn:ietf:params:xml:ns:epp-1.0 all"`
	None             string `xml:"urn:ietf:params:xml:ns:epp-1.0 none"`
	Null             string `xml:"urn:ietf:params:xml:ns:epp-1.0 null"`
	Other            string `xml:"urn:ietf:params:xml:ns:epp-1.0 other"`
	Personal         string `xml:"urn:ietf:params:xml:ns:epp-1.0 personal"`
	PersonalAndOther string `xml:"urn:ietf:params:xml:ns:epp-1.0 personalAndOther"`
}

type DcpExpiryType struct {
	Absolute time.Time `xml:"urn:ietf:params:xml:ns:epp-1.0 absolute"`
	Relative string    `xml:"urn:ietf:params:xml:ns:epp-1.0 relative"`
}

func (t *DcpExpiryType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DcpExpiryType
	var layout struct {
		*T
		Absolute *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 absolute"`
	}
	layout.T = (*T)(t)
	layout.Absolute = (*xsdDateTime)(&layout.T.Absolute)
	return e.EncodeElement(layout, start)
}
func (t *DcpExpiryType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DcpExpiryType
	var overlay struct {
		*T
		Absolute *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 absolute"`
	}
	overlay.T = (*T)(t)
	overlay.Absolute = (*xsdDateTime)(&overlay.T.Absolute)
	return d.DecodeElement(&overlay, &start)
}

type DcpOursType struct {
	RecDesc DcpRecDescType `xml:"urn:ietf:params:xml:ns:epp-1.0 recDesc,omitempty"`
}

type DcpPurposeType struct {
	Admin   string `xml:"urn:ietf:params:xml:ns:epp-1.0 admin,omitempty"`
	Contact string `xml:"urn:ietf:params:xml:ns:epp-1.0 contact,omitempty"`
	Other   string `xml:"urn:ietf:params:xml:ns:epp-1.0 other,omitempty"`
	Prov    string `xml:"urn:ietf:params:xml:ns:epp-1.0 prov,omitempty"`
}

// Must be at least 1 items long
type DcpRecDescType string

type DcpRecipientType struct {
	Other     string        `xml:"urn:ietf:params:xml:ns:epp-1.0 other,omitempty"`
	Ours      []DcpOursType `xml:"urn:ietf:params:xml:ns:epp-1.0 ours,omitempty"`
	Public    string        `xml:"urn:ietf:params:xml:ns:epp-1.0 public,omitempty"`
	Same      string        `xml:"urn:ietf:params:xml:ns:epp-1.0 same,omitempty"`
	Unrelated string        `xml:"urn:ietf:params:xml:ns:epp-1.0 unrelated,omitempty"`
}

type DcpRetentionType struct {
	Business   string `xml:"urn:ietf:params:xml:ns:epp-1.0 business"`
	Indefinite string `xml:"urn:ietf:params:xml:ns:epp-1.0 indefinite"`
	Legal      string `xml:"urn:ietf:params:xml:ns:epp-1.0 legal"`
	None       string `xml:"urn:ietf:params:xml:ns:epp-1.0 none"`
	Stated     string `xml:"urn:ietf:params:xml:ns:epp-1.0 stated"`
}

type DcpStatementType struct {
	Purpose   DcpPurposeType   `xml:"urn:ietf:params:xml:ns:epp-1.0 purpose"`
	Recipient DcpRecipientType `xml:"urn:ietf:params:xml:ns:epp-1.0 recipient"`
	Retention DcpRetentionType `xml:"urn:ietf:params:xml:ns:epp-1.0 retention"`
}

type DcpType struct {
	Access    DcpAccessType      `xml:"urn:ietf:params:xml:ns:epp-1.0 access"`
	Statement []DcpStatementType `xml:"urn:ietf:params:xml:ns:epp-1.0 statement"`
	Expiry    DcpExpiryType      `xml:"urn:ietf:params:xml:ns:epp-1.0 expiry,omitempty"`
}

type DeleteNotifyType struct {
	Delete SIDType `xml:"urn:se:iis:xml:epp:iis-1.2 delete"`
}

type DiscloseType struct {
	Name  []IntLocType `xml:"urn:ietf:params:xml:ns:contact-1.0 name,omitempty"`
	Org   []IntLocType `xml:"urn:ietf:params:xml:ns:contact-1.0 org,omitempty"`
	Addr  []IntLocType `xml:"urn:ietf:params:xml:ns:contact-1.0 addr,omitempty"`
	Voice string       `xml:"urn:ietf:params:xml:ns:contact-1.0 voice,omitempty"`
	Fax   string       `xml:"urn:ietf:params:xml:ns:contact-1.0 fax,omitempty"`
	Email string       `xml:"urn:ietf:params:xml:ns:contact-1.0 email,omitempty"`
	Flag  bool         `xml:"flag,attr"`
}

type DsDataType struct {
	KeyTag     uint        `xml:"urn:ietf:params:xml:ns:secDNS-1.1 keyTag"`
	Alg        byte        `xml:"urn:ietf:params:xml:ns:secDNS-1.1 alg"`
	DigestType byte        `xml:"urn:ietf:params:xml:ns:secDNS-1.1 digestType"`
	Digest     []byte      `xml:"urn:ietf:params:xml:ns:secDNS-1.1 digest"`
	KeyData    KeyDataType `xml:"urn:ietf:params:xml:ns:secDNS-1.1 keyData,omitempty"`
}

func (t *DsDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DsDataType
	var layout struct {
		*T
		Digest *xsdHexBinary `xml:"urn:ietf:params:xml:ns:secDNS-1.1 digest"`
	}
	layout.T = (*T)(t)
	layout.Digest = (*xsdHexBinary)(&layout.T.Digest)
	return e.EncodeElement(layout, start)
}
func (t *DsDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DsDataType
	var overlay struct {
		*T
		Digest *xsdHexBinary `xml:"urn:ietf:params:xml:ns:secDNS-1.1 digest"`
	}
	overlay.T = (*T)(t)
	overlay.Digest = (*xsdHexBinary)(&overlay.T.Digest)
	return d.DecodeElement(&overlay, &start)
}

type DsOrKeyType struct {
	MaxSigLife int           `xml:"urn:ietf:params:xml:ns:secDNS-1.1 maxSigLife,omitempty"`
	DsData     []DsDataType  `xml:"urn:ietf:params:xml:ns:secDNS-1.1 dsData"`
	KeyData    []KeyDataType `xml:"urn:ietf:params:xml:ns:secDNS-1.1 keyData"`
}

type DsType struct {
	DsData []DsDataType `xml:"urn:ietf:params:xml:ns:secDNS-1.0 dsData"`
}

// Must match the pattern (\+[0-9]{1,3}\.[0-9]{1,14})?
type E164StringType string

type E164Type struct {
	E164StringType E164StringType `xml:",chardata"`
	X              string         `xml:"x,attr,omitempty"`
}

type EppType struct {
	Greeting  *GreetingType `xml:"urn:ietf:params:xml:ns:epp-1.0 greeting,omitempty"`
	Hello     *string       `xml:"urn:ietf:params:xml:ns:epp-1.0 hello,omitempty"`
	Command   *CommandType  `xml:"urn:ietf:params:xml:ns:epp-1.0 command,omitempty"`
	Response  *ResponseType `xml:"urn:ietf:params:xml:ns:epp-1.0 response,omitempty"`
	Extension *ExtAnyType   `xml:"urn:ietf:params:xml:ns:epp-1.0 extension,omitempty"`
}

type ErrValueType struct {
	Item string `xml:",any"`
}

type ExtAnyType []string

func (a ExtAnyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:" item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *ExtAnyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

type ExtAuthInfoType struct {
	Item string `xml:",any"`
}

type ExtErrValueType struct {
	Value  ErrValueType `xml:"urn:ietf:params:xml:ns:epp-1.0 value"`
	Reason MsgType      `xml:"urn:ietf:params:xml:ns:epp-1.0 reason"`
}

type ExtURIType struct {
	ExtURI []string `xml:"urn:ietf:params:xml:ns:epp-1.0 extURI"`
}

type GreetingType struct {
	SvID    SIDType     `xml:"urn:ietf:params:xml:ns:epp-1.0 svID"`
	SvDate  time.Time   `xml:"urn:ietf:params:xml:ns:epp-1.0 svDate"`
	SvcMenu SvcMenuType `xml:"urn:ietf:params:xml:ns:epp-1.0 svcMenu"`
	Dcp     DcpType     `xml:"urn:ietf:params:xml:ns:epp-1.0 dcp"`
}

func (t *GreetingType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T GreetingType
	var layout struct {
		*T
		SvDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 svDate"`
	}
	layout.T = (*T)(t)
	layout.SvDate = (*xsdDateTime)(&layout.T.SvDate)
	return e.EncodeElement(layout, start)
}
func (t *GreetingType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T GreetingType
	var overlay struct {
		*T
		SvDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 svDate"`
	}
	overlay.T = (*T)(t)
	overlay.SvDate = (*xsdDateTime)(&overlay.T.SvDate)
	return d.DecodeElement(&overlay, &start)
}

type HostAttrType struct {
	HostName LabelType  `xml:"urn:ietf:params:xml:ns:domain-1.0 hostName"`
	HostAddr []AddrType `xml:"urn:ietf:params:xml:ns:domain-1.0 hostAddr,omitempty"`
}

// May be one of all, del, none, sub
type HostsType string

type InfDataType struct {
	Id         ClIDType         `xml:"urn:ietf:params:xml:ns:contact-1.0 id"`
	Roid       RoidType         `xml:"urn:ietf:params:xml:ns:contact-1.0 roid"`
	Status     []StatusType     `xml:"urn:ietf:params:xml:ns:contact-1.0 status"`
	PostalInfo []PostalInfoType `xml:"urn:ietf:params:xml:ns:contact-1.0 postalInfo"`
	Voice      E164Type         `xml:"urn:ietf:params:xml:ns:contact-1.0 voice,omitempty"`
	Fax        E164Type         `xml:"urn:ietf:params:xml:ns:contact-1.0 fax,omitempty"`
	Email      MinTokenType     `xml:"urn:ietf:params:xml:ns:contact-1.0 email"`
	ClID       ClIDType         `xml:"urn:ietf:params:xml:ns:contact-1.0 clID"`
	CrID       ClIDType         `xml:"urn:ietf:params:xml:ns:contact-1.0 crID"`
	CrDate     time.Time        `xml:"urn:ietf:params:xml:ns:contact-1.0 crDate"`
	UpID       ClIDType         `xml:"urn:ietf:params:xml:ns:contact-1.0 upID,omitempty"`
	UpDate     time.Time        `xml:"urn:ietf:params:xml:ns:contact-1.0 upDate,omitempty"`
	TrDate     time.Time        `xml:"urn:ietf:params:xml:ns:contact-1.0 trDate,omitempty"`
	AuthInfo   AuthInfoType     `xml:"urn:ietf:params:xml:ns:contact-1.0 authInfo,omitempty"`
	Disclose   DiscloseType     `xml:"urn:ietf:params:xml:ns:contact-1.0 disclose,omitempty"`
}

func (t *InfDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T InfDataType
	var layout struct {
		*T
		CrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 crDate"`
		UpDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 upDate,omitempty"`
		TrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 trDate,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CrDate = (*xsdDateTime)(&layout.T.CrDate)
	layout.UpDate = (*xsdDateTime)(&layout.T.UpDate)
	layout.TrDate = (*xsdDateTime)(&layout.T.TrDate)
	return e.EncodeElement(layout, start)
}
func (t *InfDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InfDataType
	var overlay struct {
		*T
		CrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 crDate"`
		UpDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 upDate,omitempty"`
		TrDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 trDate,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CrDate = (*xsdDateTime)(&overlay.T.CrDate)
	overlay.UpDate = (*xsdDateTime)(&overlay.T.UpDate)
	overlay.TrDate = (*xsdDateTime)(&overlay.T.TrDate)
	return d.DecodeElement(&overlay, &start)
}

type InfoNameType struct {
	LabelType LabelType `xml:",chardata"`
	Hosts     HostsType `xml:"hosts,attr,omitempty"`
}

func (t *InfoNameType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T InfoNameType
	var overlay struct {
		*T
		Hosts *HostsType `xml:"hosts,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Hosts = (*HostsType)(&overlay.T.Hosts)
	return d.DecodeElement(&overlay, &start)
}

type InfoType struct {
	Name     InfoNameType `xml:"urn:ietf:params:xml:ns:domain-1.0 name"`
	AuthInfo AuthInfoType `xml:"urn:ietf:params:xml:ns:domain-1.0 authInfo,omitempty"`
}

type IntLocType struct {
	Type PostalInfoEnumType `xml:"type,attr"`
}

// May be one of v4, v6
type IpType string

type KeyDataType struct {
	Flags    uint    `xml:"urn:ietf:params:xml:ns:secDNS-1.1 flags"`
	Protocol byte    `xml:"urn:ietf:params:xml:ns:secDNS-1.1 protocol"`
	Alg      byte    `xml:"urn:ietf:params:xml:ns:secDNS-1.1 alg"`
	PubKey   KeyType `xml:"urn:ietf:params:xml:ns:secDNS-1.1 pubKey"`
}

type KeyType []byte

func (t *KeyType) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t KeyType) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type LabelType string

type LoginSvcType struct {
	ObjURI       []string   `xml:"urn:ietf:params:xml:ns:epp-1.0 objURI"`
	SvcExtension ExtURIType `xml:"urn:ietf:params:xml:ns:epp-1.0 svcExtension,omitempty"`
}

type LoginType struct {
	ClID    ClIDType         `xml:"urn:ietf:params:xml:ns:epp-1.0 clID"`
	Pw      PwType           `xml:"urn:ietf:params:xml:ns:epp-1.0 pw"`
	NewPW   PwType           `xml:"urn:ietf:params:xml:ns:epp-1.0 newPW,omitempty"`
	Options CredsOptionsType `xml:"urn:ietf:params:xml:ns:epp-1.0 options"`
	Svcs    LoginSvcType     `xml:"urn:ietf:params:xml:ns:epp-1.0 svcs"`
}

type MIDType struct {
	Id []ClIDType `xml:"urn:ietf:params:xml:ns:contact-1.0 id"`
}

type MNameType struct {
	Name []LabelType `xml:"urn:ietf:params:xml:ns:host-1.0 name"`
}

// Must be at least 1 items long
type MinTokenType string

type MixedMsgType struct {
	Items []string `xml:",any"`
	Lang  string   `xml:"lang,attr,omitempty"`
}

func (t *MixedMsgType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MixedMsgType
	var overlay struct {
		*T
		Lang *string `xml:"lang,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Lang = (*string)(&overlay.T.Lang)
	return d.DecodeElement(&overlay, &start)
}

type MsgQType struct {
	QDate time.Time    `xml:"urn:ietf:params:xml:ns:epp-1.0 qDate,omitempty"`
	Msg   MixedMsgType `xml:"urn:ietf:params:xml:ns:epp-1.0 msg,omitempty"`
	Count uint64       `xml:"count,attr"`
}

func (t *MsgQType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MsgQType
	var layout struct {
		*T
		QDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 qDate,omitempty"`
	}
	layout.T = (*T)(t)
	layout.QDate = (*xsdDateTime)(&layout.T.QDate)
	return e.EncodeElement(layout, start)
}
func (t *MsgQType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MsgQType
	var overlay struct {
		*T
		QDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:epp-1.0 qDate,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.QDate = (*xsdDateTime)(&overlay.T.QDate)
	return d.DecodeElement(&overlay, &start)
}

type MsgType struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"lang,attr,omitempty"`
}

func (t *MsgType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MsgType
	var overlay struct {
		*T
		Lang *string `xml:"lang,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Lang = (*string)(&overlay.T.Lang)
	return d.DecodeElement(&overlay, &start)
}

type NsType struct {
	HostObj []LabelType `xml:"urn:se:iis:xml:epp:iis-1.2 hostObj"`
}

// Must match the pattern \[[A-Z][A-Z]\].{1,123}
type OrgnoType string

// May be one of y, m
type PUnitType string

type PaCLIDType struct {
	ClIDType ClIDType `xml:",chardata"`
	PaResult bool     `xml:"paResult,attr"`
}

type PaNameType struct {
	LabelType LabelType `xml:",chardata"`
	PaResult  bool      `xml:"paResult,attr"`
}

type PanDataType struct {
	Name   PaNameType `xml:"urn:ietf:params:xml:ns:host-1.0 name"`
	PaTRID TrIDType   `xml:"urn:ietf:params:xml:ns:host-1.0 paTRID"`
	PaDate time.Time  `xml:"urn:ietf:params:xml:ns:host-1.0 paDate"`
}

func (t *PanDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PanDataType
	var layout struct {
		*T
		PaDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:host-1.0 paDate"`
	}
	layout.T = (*T)(t)
	layout.PaDate = (*xsdDateTime)(&layout.T.PaDate)
	return e.EncodeElement(layout, start)
}
func (t *PanDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PanDataType
	var overlay struct {
		*T
		PaDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:host-1.0 paDate"`
	}
	overlay.T = (*T)(t)
	overlay.PaDate = (*xsdDateTime)(&overlay.T.PaDate)
	return d.DecodeElement(&overlay, &start)
}

type PeriodType struct {
	Value uint      `xml:",chardata"`
	Unit  PUnitType `xml:"unit,attr"`
}

// May be one of ack, req
type PollOpType string

type PollType struct {
	Op    PollOpType `xml:"op,attr"`
	MsgID string     `xml:"msgID,attr,omitempty"`
}

// May be one of loc, int
type PostalInfoEnumType string

type PostalInfoType struct {
	Name PostalLineType     `xml:"urn:ietf:params:xml:ns:contact-1.0 name"`
	Org  string             `xml:"urn:ietf:params:xml:ns:contact-1.0 org,omitempty"`
	Addr AddrType           `xml:"urn:ietf:params:xml:ns:contact-1.0 addr"`
	Type PostalInfoEnumType `xml:"type,attr"`
}

// Must be at least 1 items long
type PostalLineType string

type PwAuthInfoType struct {
	Value string   `xml:",chardata"`
	Roid  RoidType `xml:"roid,attr,omitempty"`
}

// Must be at least 6 items long
type PwType string

type ReadWriteType struct {
	Item string `xml:",any"`
}

// Must be at least 1 items long
type ReasonBaseType string

type ReasonType struct {
	ReasonBaseType ReasonBaseType `xml:",chardata"`
	Lang           string         `xml:"lang,attr,omitempty"`
}

type RemType struct {
	All     bool          `xml:"urn:ietf:params:xml:ns:secDNS-1.1 all"`
	DsData  []DsDataType  `xml:"urn:ietf:params:xml:ns:secDNS-1.1 dsData"`
	KeyData []KeyDataType `xml:"urn:ietf:params:xml:ns:secDNS-1.1 keyData"`
}

type RenDataType struct {
	Name   LabelType `xml:"urn:ietf:params:xml:ns:domain-1.0 name"`
	ExDate time.Time `xml:"urn:ietf:params:xml:ns:domain-1.0 exDate,omitempty"`
}

func (t *RenDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RenDataType
	var layout struct {
		*T
		ExDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:domain-1.0 exDate,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ExDate = (*xsdDateTime)(&layout.T.ExDate)
	return e.EncodeElement(layout, start)
}
func (t *RenDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RenDataType
	var overlay struct {
		*T
		ExDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:domain-1.0 exDate,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ExDate = (*xsdDateTime)(&overlay.T.ExDate)
	return d.DecodeElement(&overlay, &start)
}

type RenewType struct {
	Name       LabelType  `xml:"urn:ietf:params:xml:ns:domain-1.0 name"`
	CurExpDate time.Time  `xml:"urn:ietf:params:xml:ns:domain-1.0 curExpDate"`
	Period     PeriodType `xml:"urn:ietf:params:xml:ns:domain-1.0 period,omitempty"`
}

func (t *RenewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RenewType
	var layout struct {
		*T
		CurExpDate *xsdDate `xml:"urn:ietf:params:xml:ns:domain-1.0 curExpDate"`
	}
	layout.T = (*T)(t)
	layout.CurExpDate = (*xsdDate)(&layout.T.CurExpDate)
	return e.EncodeElement(layout, start)
}
func (t *RenewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RenewType
	var overlay struct {
		*T
		CurExpDate *xsdDate `xml:"urn:ietf:params:xml:ns:domain-1.0 curExpDate"`
	}
	overlay.T = (*T)(t)
	overlay.CurExpDate = (*xsdDate)(&overlay.T.CurExpDate)
	return d.DecodeElement(&overlay, &start)
}

type ResponseType struct {
	Result    []ResultType `xml:"urn:ietf:params:xml:ns:epp-1.0 result"`
	MsgQ      MsgQType     `xml:"urn:ietf:params:xml:ns:epp-1.0 msgQ,omitempty"`
	ResData   ExtAnyType   `xml:"urn:ietf:params:xml:ns:epp-1.0 resData,omitempty"`
	Extension ExtAnyType   `xml:"urn:ietf:params:xml:ns:epp-1.0 extension,omitempty"`
	TrID      TrIDType     `xml:"urn:ietf:params:xml:ns:epp-1.0 trID"`
}

// May be one of 1000, 1001, 1300, 1301, 1500, 2000, 2001, 2002, 2003, 2004, 2005, 2100, 2101, 2102, 2103, 2104, 2105, 2106, 2200, 2201, 2202, 2300, 2301, 2302, 2303, 2304, 2305, 2306, 2307, 2308, 2400, 2500, 2501, 2502
type ResultCodeType uint

type ResultType struct {
	Msg      MsgType         `xml:"urn:ietf:params:xml:ns:epp-1.0 msg"`
	Value    ErrValueType    `xml:"urn:ietf:params:xml:ns:epp-1.0 value"`
	ExtValue ExtErrValueType `xml:"urn:ietf:params:xml:ns:epp-1.0 extValue"`
	Code     ResultCodeType  `xml:"code,attr"`
}

// Must match the pattern (\w|_){1,80}-\w{1,8}
type RoidType string

type SIDType struct {
	Id ClIDType `xml:"urn:ietf:params:xml:ns:contact-1.0 id"`
}

type SNameType struct {
	Name LabelType `xml:"urn:ietf:params:xml:ns:host-1.0 name"`
}

type StatusType struct {
	Value string          `xml:",chardata"`
	S     StatusValueType `xml:"s,attr"`
	Lang  string          `xml:"lang,attr,omitempty"`
}

func (t *StatusType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T StatusType
	var overlay struct {
		*T
		Lang *string `xml:"lang,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Lang = (*string)(&overlay.T.Lang)
	return d.DecodeElement(&overlay, &start)
}

// May be one of clientDeleteProhibited, clientTransferProhibited, clientUpdateProhibited, linked, ok, pendingCreate, pendingDelete, pendingTransfer, pendingUpdate, serverDeleteProhibited, serverTransferProhibited, serverUpdateProhibited
type StatusValueType string

type SvcMenuType struct {
	Version      []VersionType `xml:"urn:ietf:params:xml:ns:epp-1.0 version"`
	Lang         []string      `xml:"urn:ietf:params:xml:ns:epp-1.0 lang"`
	ObjURI       []string      `xml:"urn:ietf:params:xml:ns:epp-1.0 objURI"`
	SvcExtension ExtURIType    `xml:"urn:ietf:params:xml:ns:epp-1.0 svcExtension,omitempty"`
}

// Must be at least 3 items long
type TrIDStringType string

type TrIDType struct {
	ClTRID TrIDStringType `xml:"urn:ietf:params:xml:ns:epp-1.0 clTRID,omitempty"`
	SvTRID TrIDStringType `xml:"urn:ietf:params:xml:ns:epp-1.0 svTRID"`
}

// May be one of clientApproved, clientCancelled, clientRejected, pending, serverApproved, serverCancelled
type TrStatusType string

type TransferNotifyType struct {
	TrnData TrnDataType `xml:"urn:se:iis:xml:epp:iis-1.2 trnData"`
}

// May be one of approve, cancel, query, reject, request
type TransferOpType string

type TransferType struct {
	Ns NsType `xml:"urn:se:iis:xml:epp:iis-1.2 ns"`
}

type TrnDataType struct {
	Id       ClIDType     `xml:"urn:ietf:params:xml:ns:contact-1.0 id"`
	TrStatus TrStatusType `xml:"urn:ietf:params:xml:ns:contact-1.0 trStatus"`
	ReID     ClIDType     `xml:"urn:ietf:params:xml:ns:contact-1.0 reID"`
	ReDate   time.Time    `xml:"urn:ietf:params:xml:ns:contact-1.0 reDate"`
	AcID     ClIDType     `xml:"urn:ietf:params:xml:ns:contact-1.0 acID"`
	AcDate   time.Time    `xml:"urn:ietf:params:xml:ns:contact-1.0 acDate"`
}

func (t *TrnDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T TrnDataType
	var layout struct {
		*T
		ReDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 reDate"`
		AcDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 acDate"`
	}
	layout.T = (*T)(t)
	layout.ReDate = (*xsdDateTime)(&layout.T.ReDate)
	layout.AcDate = (*xsdDateTime)(&layout.T.AcDate)
	return e.EncodeElement(layout, start)
}
func (t *TrnDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T TrnDataType
	var overlay struct {
		*T
		ReDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 reDate"`
		AcDate *xsdDateTime `xml:"urn:ietf:params:xml:ns:contact-1.0 acDate"`
	}
	overlay.T = (*T)(t)
	overlay.ReDate = (*xsdDateTime)(&overlay.T.ReDate)
	overlay.AcDate = (*xsdDateTime)(&overlay.T.AcDate)
	return d.DecodeElement(&overlay, &start)
}

type UpdateNotifyType struct {
	InfData InfDataType `xml:"urn:se:iis:xml:epp:iis-1.2 infData"`
}

type UpdateType struct {
	Rem    RemType     `xml:"urn:ietf:params:xml:ns:secDNS-1.1 rem,omitempty"`
	Add    DsOrKeyType `xml:"urn:ietf:params:xml:ns:secDNS-1.1 add,omitempty"`
	Chg    ChgType     `xml:"urn:ietf:params:xml:ns:secDNS-1.1 chg,omitempty"`
	Urgent bool        `xml:"urgent,attr,omitempty"`
}

func (t *UpdateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UpdateType
	var overlay struct {
		*T
		Urgent *bool `xml:"urgent,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Urgent = (*bool)(&overlay.T.Urgent)
	return d.DecodeElement(&overlay, &start)
}

// May be one of 1.0
type VersionType string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdHexBinary []byte

func (b *xsdHexBinary) UnmarshalText(text []byte) (err error) {
	*b, err = hex.DecodeString(string(text))
	return
}
func (b xsdHexBinary) MarshalText() ([]byte, error) {
	n := hex.EncodedLen(len(b))
	buf := make([]byte, n)
	hex.Encode(buf, []byte(b))
	return buf, nil
}
