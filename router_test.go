package eppserver

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"aqwari.net/xml/xmltree"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_buildPath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "ack-poll.xml",
			want:  "command>poll",
		},
		{
			input: "check-contact.xml",
			want:  "command>check>contact",
		},
		{
			input: "check-domain.xml",
			want:  "command>check>contact",
		},
		{
			input: "login.xml",
			want:  "command>login",
		},
		{
			input: "info-domain.xml",
			want:  "command>info>domain",
		},
		{
			input: "logout.xml",
			want:  "command>logout",
		},
		{
			input: "hello.xml",
			want:  "hello",
		},
		{
			input: "transfer-domain.xml",
			want:  "command>transfer>domain",
		},
	}

	for _, tt := range tests {
		fileData, err := ioutil.ReadFile(filepath.Join("xml", "commands", tt.input))
		require.Nil(t, err)

		t.Run(tt.input, func(t *testing.T) {
			root, err := xmltree.Parse(fileData)
			require.Nil(t, err)

			path, err := buildPath(root)
			require.Nil(t, err)

			assert.Equal(t, tt.want, path)
		})
	}
}
