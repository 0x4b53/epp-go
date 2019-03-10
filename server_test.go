package epp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	cert := generateCertificate()

	srv := Server{
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		Handler: func(s *Session, in []byte) ([]byte, error) {
			data := fmt.Sprintf("response: %s", string(in))
			return []byte(data), nil
		},
		Greeting: func(s *Session) ([]byte, error) {
			return []byte("hello"), nil
		},
		Addr: ":9889",
	}
	defer srv.Stop()

	didStart := make(chan struct{})
	srv.OnStarted(func() {
		didStart <- struct{}{}
	})

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	<-didStart
	client := &Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	greeting, err := client.Connect(":9889")
	require.Nil(t, err)

	assert.Equal(t, string(greeting), "hello")

	for i := 0; i < 5; i++ {
		data := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
		<epp xmlns="urn:ietf:params:xml:ns:epp-1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd">
		  <hello>message %d</hello>
		</epp>`, i)
		response, err := client.Send([]byte(data))
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("response: %s", data), string(response))
	}
}

func generateCertificate() tls.Certificate {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			CommonName:   "epp.example.test",
			Organization: []string{"Simple Server Test"},
			Country:      []string{"SE"},
			Locality:     []string{"Stockholm"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 1),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	certificate, _ := x509.CreateCertificate(rand.Reader, cert, cert, key.Public(), key)

	return tls.Certificate{
		Certificate: [][]byte{certificate},
		PrivateKey:  key,
	}
}
