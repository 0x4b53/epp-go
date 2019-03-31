package epp

import (
	"encoding/xml"
	"fmt"
	"net"
	"testing"

	"github.com/bombsimon/epp-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadWriteMessage(t *testing.T) {
	conn1, conn2 := net.Pipe()

	go func() {
		for i := 0; i < 10; i++ {
			err := WriteMessage(conn1, []byte(fmt.Sprintf("ping %d", i)))
			require.Nil(t, err)

			message, err := ReadMessage(conn1)
			require.Nil(t, err)
			assert.Equal(t, fmt.Sprintf("pong %d", i), string(message))
		}
	}()

	for i := 0; i < 10; i++ {
		message, err := ReadMessage(conn2)
		require.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("ping %d", i), string(message))

		err = WriteMessage(conn2, []byte(fmt.Sprintf("pong %d", i)))
		require.Nil(t, err)
	}
}

func TestEncode(t *testing.T) {
	dc := types.DomainCreate{
		Name: "example.net",
		Period: types.Period{
			Value: 12,
			Unit:  "m",
		},
		NameServer: types.NameServer{
			HostObject: []string{
				"ns1.example.net",
				"ns2.example.net",
			},
		},
		Registrant: "registrant-00001",
		Contacts: []types.Contact{
			{
				Name: "contact-00001",
				Type: "tech",
			},
			{
				Name: "contact-00002",
				Type: "admin",
			},
		},
		AuthInfo: types.AuthInfo{
			Password: "some-password",
		},
	}

	expectedXMLWithNS := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <command>
    <create>
      <domain:create xmlns:domain="urn:ietf:params:xml:ns:domain-1.0">
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
      </domain:create>
    </create>
  </command>
</epp>
`)

	encoded, err := Encode(dc, ClientXMLAttributes())

	require.Nil(t, err)
	assert.Equal(t, expectedXMLWithNS, encoded)
}

func TestDecode(t *testing.T) {
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

	dc := types.DomainCreate{}
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

func ExampleAddNamespace() {
	// Construct the response with basic data.
	diResponse := types.DomainInfoDataType{
		InfoData: types.DomainInfoData{
			Name: "example.se",
			ROID: "DOMAIN_0000000000-SE",
			Status: []types.DomainStatus{
				{
					DomainStatusType: types.DomainStatusOk,
				},
			},
			Host: []string{
				"ns1.example.se", "ns2.example.se",
			},
			ClientID: "Some Client",
			CreateID: "Some Client",
			UpdateID: "Some Client",
		},
	}

	// Add extension data from extension iis-1.2.
	diIISExtensionResponse := types.IISExtensionInfoDataType{
		InfoData: types.IISExtensionInfoData{
			State:        "active",
			ClientDelete: false,
		},
	}

	// ADd extension data from secDNS-1.1.
	diDNSSECExtensionResponse := types.DNSSECExtensionInfoDataType{
		InfoData: types.DNSSECOrKeyData{
			DNSSECData: []types.DNSSEC{
				{
					KeyTag:     195550,
					Algorithm:  3,
					DigestType: 5,
					Digest:     "FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102",
				},
			},
		},
	}

	response := types.Response{
		Result: []types.Result{
			{
				Code:    EppOk.Code(),
				Message: EppOk.Message(),
			},
		},
		ResultData: diResponse,
		Extension: struct {
			types.IISExtensionInfoDataType
			types.DNSSECExtensionInfoDataType
		}{
			diIISExtensionResponse,
			diDNSSECExtensionResponse,
		},
	}

	b, _ := Encode(response, ClientXMLAttributes())

	fmt.Println(string(b))

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
	//   <response>
	//     <result code="1000">
	//       <msg>Command completed successfully</msg>
	//     </result>
	//     <resData>
	//       <domain:infData xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xmlns="urn:ietf:params:xml:ns:domain-1.0">
	//         <domain:name>example.se</domain:name>
	//         <domain:roid>DOMAIN_0000000000-SE</domain:roid>
	//         <domain:status s="ok" />
	//         <domain:host>ns1.example.se</domain:host>
	//         <domain:host>ns2.example.se</domain:host>
	//         <domain:clID>Some Client</domain:clID>
	//         <domain:crID>Some Client</domain:crID>
	//         <domain:upID>Some Client</domain:upID>
	//         <domain:authInfo />
	//       </domain:infData>
	//     </resData>
	//     <extension>
	//       <iis:infData xmlns:iis="urn:se:iis:xml:epp:iis-1.2" xmlns="urn:se:iis:xml:epp:iis-1.2">
	//         <iis:state>active</iis:state>
	//       </iis:infData>
	//       <sec:infData xmlns:sec="urn:ietf:params:xml:ns:secDNS-1.1" xmlns="urn:ietf:params:xml:ns:secDNS-1.1">
	//         <sec:dsData>
	//           <sec:keyTag>195550</sec:keyTag>
	//           <sec:alg>3</sec:alg>
	//           <sec:digestType>5</sec:digestType>
	//           <sec:digest>FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102</sec:digest>
	//         </sec:dsData>
	//       </sec:infData>
	//     </extension>
	//     <trID>
	//       <clTRID />
	//       <svTRID />
	//     </trID>
	//   </response>
	// </epp>
}
