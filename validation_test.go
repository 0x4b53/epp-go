package epp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	xsd "github.com/lestrrat-go/libxml2/xsd"
)

func TestValidator_setupSchema(t *testing.T) {
	validator, err := NewValidator("xml/index.xsd")

	require.Nil(t, err)
	require.NotNil(t, validator)

	cases := []struct {
		description string
		xml         []byte
		errContains string
		xmlErrors   []string
	}{
		{
			description: "no xml should not be valid",
			xml:         []byte{},
			errContains: "failed to create parse context",
		},
		{
			description: "namespace for EPP tag is requried",
			xml:         []byte(`<epp><command></command></epp>`),
			errContains: "schema validation failed",
			xmlErrors:   []string{"Element 'epp': No matching global declaration available for the validation root."},
		},
		{
			description: "attributes are verified",
			xml: []byte(`
				<epp xmlns="urn:ietf:params:xml:ns:epp-1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
				  <command>
					<poll msgID="3" op="-INVALID-"/>
					<clTRID>ABC-12345</clTRID>
				  </command>
				</epp>`),
			errContains: "schema validation failed",
			xmlErrors:   []string{"'-INVALID-' is not an element of the set {'ack', 'req'}", "'-INVALID-' is not a valid value of the atomic type"},
		},
		{
			description: "valid XML, including type ns",
			xml: []byte(`
				<epp xmlns="urn:ietf:params:xml:ns:epp-1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
				  <command>
					<check>
					  <domain:check xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xsi:schemaLocation="urn:ietf:params:xml:ns:domain-1.0  domain-1.0.xsd">
						<domain:name>example1.se</domain:name>
						<domain:name>example2.se</domain:name>
					  </domain:check>
					</check>
					<clTRID>ABC-12345</clTRID>
				  </command>
				</epp>`),
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			err := validator.Validate(tc.xml)

			if tc.errContains == "" {
				require.Nil(t, err)

				return
			}

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), tc.errContains)

			xErr, ok := err.(xsd.SchemaValidationError)
			if !ok {
				return
			}

			xErrors := xErr.Errors()
			if len(xErrors) != len(tc.xmlErrors) {
				t.Logf("all errors not caught, got %d errors:\n", len(xErrors))

				for i, err := range xErrors {
					t.Logf("%-3d %s\n", i, err.Error())
				}

				assert.Fail(t, "all errors are not caught")
				return
			}

			for i, err := range tc.xmlErrors {
				assert.Contains(t, xErrors[i].Error(), err)
			}
		})
	}
}
