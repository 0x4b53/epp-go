package types

import (
	"encoding/xml"
	"fmt"
	"testing"

	"aqwari.net/xml/xmltree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	dc := DomainCreate{
		Name: "example.net",
		Period: Period{
			Value: 12,
			Unit:  "m",
		},
		NameServer: NameServer{
			HostObject: []string{
				"ns1.example.net",
				"ns2.example.net",
			},
		},
		Registrant: "registrant-00001",
		Contacts: []Contact{
			{
				Name: "contact-00001",
				Type: "tech",
			},
			{
				Name: "contact-00002",
				Type: "admin",
			},
		},
		AuthInfo: AuthInfo{
			Password: "some-password",
		},
	}

	expectedXML := []byte(`<DomainCreate>
  <command>
    <create>
      <create>
        <name>example.net</name>
        <period unit="m">12</period>
        <ns>
          <hostObj>ns1.example.net</hostObj>
          <hostObj>ns2.example.net</hostObj>
        </ns>
        <registrant>registrant-00001</registrant>
        <contact type="tech">contact-00001</contact>
        <contact type="admin">contact-00002</contact>
        <authInfo>
          <pw>some-password</pw>
        </authInfo>
      </create>
    </create>
  </command>
</DomainCreate>`)

	expectewdXMLWithNS := []byte(`<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <command>
    <create>
      <create xmlns:domain="urn:ietf:params:xml:ns:domain-1.0 domain-1.0">
        <domain:name>
          example.net
        </domain:name>
        <domain:period unit="m">
          12
        </domain:period>
        <domain:ns>
          <domain:hostObj>
            ns1.example.net
          </domain:hostObj>
          <domain:hostObj>
            ns2.example.net
          </domain:hostObj>
        </domain:ns>
        <domain:registrant>
          registrant-00001
        </domain:registrant>
        <domain:contact type="tech">
          contact-00001
        </domain:contact>
        <domain:contact type="admin">
          contact-00002
        </domain:contact>
        <domain:authInfo>
          <domain:pw>
            some-password
          </domain:pw>
        </domain:authInfo>
      </create>
    </create>
  </command>
</epp>
`)

	// Ensure we can marshal the struct as is without namespaces.
	b, err := xml.MarshalIndent(dc, "", "  ")
	require.Nil(t, err)
	assert.Equal(t, expectedXML, b)

	// Parse the marshalled XML with xmltree to generate elements for each node.
	root, err := xmltree.Parse(b)
	require.Nil(t, err)

	// Overwrite the start element to set to an EPP tag with proper attributes.
	root.StartElement = xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "epp",
		},
		Attr: []xml.Attr{
			{
				Name: xml.Name{
					Space: "",
					Local: "xmlns",
				},
				Value: "urn:ietf:params:xml:ns:epp-1.0",
			},
		},
	}

	// Find all create tags
	creates := root.Search("", "create")

	// Add the custom name space `domain` to the domain create tag.
	creates[0].Children[0].SetAttr("", "xmlns:domain", "urn:ietf:params:xml:ns:domain-1.0 domain-1.0")

	// Add the name space to each tag recursively.
	addNS(creates[1].Children, "domain")

	// Ensure we can marshal after we've added the name spaces.
	bWithNS := xmltree.MarshalIndent(root, "", "  ")

	assert.Equal(t, expectewdXMLWithNS, bWithNS)
}

func addNS(children []xmltree.Element, space string) {
	for i, element := range children {
		// TODO: Why can we not set element.Name.Space?
		element.Name.Local = fmt.Sprintf("%s:%s", space, element.Name.Local)
		children[i] = element

		if len(element.Children) > 0 {
			addNS(element.Children, space)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	x := []byte(`<?xml version="1.0"?>
<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <command>
    <create>
      <domain:create xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
        <domain:name>example.net</domain:name>
        <domain:period unit="m">12</domain:period>
        <domain:ns>
          <domain:hostObj>ns1.example.net</domain:hostObj>
          <domain:hostObj>ns2.example.net</domain:hostObj>
        </domain:ns>
        <domain:registrant>registrant-00001</domain:registrant>
        <domain:contact type="tech">contact-00001</domain:contact>
        <domain:contact type="admin">contact-00002</domain:contact>
        <domain:authInfo>
          <domain:pw>some-password</domain:pw>
        </domain:authInfo>
      </domain:create>
    </create>
    <clTRID>ABC-12345</clTRID>
  </command>
</epp>`)

	dc := DomainCreate{}
	if err := xml.Unmarshal(x, &dc); err != nil {
		assert.Fail(t, "could not unmarshal")
	}

	assert.Equal(t, "example.net", dc.Name, "domain name found")
	assert.Equal(t, 12, dc.Period.Value, "period found")
	assert.Equal(t, "m", dc.Period.Unit, "period unit found")
	assert.Equal(t, "ns1.example.net", dc.NameServer.HostObject[0], "host object found")
	assert.Equal(t, "ns2.example.net", dc.NameServer.HostObject[1], "host object found")
	assert.Equal(t, "registrant-00001", dc.Registrant, "registrant found")
	assert.Equal(t, "contact-00001", dc.Contacts[0].Name, "contact found")
	assert.Equal(t, "tech", dc.Contacts[0].Type, "contact type found")
	assert.Equal(t, "contact-00002", dc.Contacts[1].Name, "contact found")
	assert.Equal(t, "admin", dc.Contacts[1].Type, "contact type found")
	assert.Equal(t, "some-password", dc.AuthInfo.Password, "auth info found")
}
