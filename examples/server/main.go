package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/xml"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

	epp "github.com/bombsimon/epp-go"
	"github.com/bombsimon/epp-go/types"
)

func main() {
	mux := epp.NewMux()

	server := epp.Server{
		IdleTimeout:      5 * time.Minute,
		MaxSessionLength: 10 * time.Minute,
		Addr:             ":4701",
		Handler:          mux.Handle,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{generateCertificate()},
			ClientAuth:   tls.RequireAnyClientCert,
		},
		Greeting: greeting,
	}

	mux.AddHandler("command/login", login)
	mux.AddHandler("command/info/domain", infoDomainWithExtension)
	mux.AddHandler("command/create/domain", createDomain)
	mux.AddHandler("command/create/contact", createContactWithExtension)

	// Support graceful shutdown.
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		server.Stop()
	}()

	log.Println("Running server...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func greeting(s *epp.Session) ([]byte, error) {
	err := verifyClientCertificate(s.ConnectionState().PeerCertificates)
	if err != nil {
		_ = s.Close()

		fmt.Println("could not verify peer certificates")

		return nil, errors.New("could not verify certificates")
	}

	greeting := types.EPPGreeting{
		Greeting: types.Greeting{
			ServerID:   "default-server",
			ServerDate: time.Now(),
			ServiceMenu: types.ServiceMenu{
				Version:  []string{"1.0"},
				Language: []string{"en"},
				ObjectURI: []string{
					types.NameSpaceDomain,
					types.NameSpaceContact,
					types.NameSpaceHost,
				},
			},
			DCP: types.DCP{
				Access: types.DCPAccess{
					All: &types.EmptyTag{},
				},
				Statement: types.DCPStatement{
					Purpose: types.DCPPurpose{
						Prov: &types.EmptyTag{},
					},
					Recipient: types.DCPRecipient{
						Ours: []types.DCPOurs{
							{},
						},
						Public: &types.EmptyTag{},
					},
					Retention: types.DCPRetention{
						Stated: &types.EmptyTag{},
					},
				},
			},
		},
	}

	return epp.Encode(greeting, epp.ServerXMLAttributes())
}

func login(s *epp.Session, data []byte) ([]byte, error) {
	domainInfo := types.DomainInfoType{
		Info: types.DomainInfo{
			Name: types.DomainInfoName{
				Name:  "example.se",
				Hosts: types.DomainHostsAll,
			},
		},
	}

	bytes, err := epp.Encode(
		domainInfo,
		epp.ClientXMLAttributes(),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	login := types.Login{}

	if err := xml.Unmarshal(data, &login); err != nil {
		return nil, err
	}

	// Authenticate the user found in login type.

	response := types.Response{
		Result: []types.Result{
			{
				Code:    epp.EppOk.Code(),
				Message: epp.EppOk.Message(),
			},
		},
		TransactionID: types.TransactionID{
			ServerTransactionID: "ABC-123",
		},
	}

	return epp.Encode(
		response,
		epp.ServerXMLAttributes(),
	)
}

func infoDomainWithExtension(s *epp.Session, data []byte) ([]byte, error) {
	di := types.DomainInfoTypeIn{}

	if err := xml.Unmarshal(data, &di); err != nil {
		return nil, err
	}

	// Assume the domain was found.

	// Construct the response with basic data.
	diResponse := types.DomainInfoDataType{
		InfoData: types.DomainInfoData{
			Name: di.Info.Name.Name,
			ROID: "DOMAIN_0000000000-SE",
			Status: []types.DomainStatus{
				{
					DomainStatusType: types.DomainStatusOk,
				},
			},
			Host: []string{
				fmt.Sprintf("ns1.%s", di.Info.Name.Name),
				fmt.Sprintf("ns2.%s", di.Info.Name.Name),
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

	// Add extension data from secDNS-1.1.
	diDNSSECExtensionResponse := types.DNSSECExtensionInfoDataType{
		InfoData: types.DNSSECOrKeyData{
			DNSSECData: []types.DNSSEC{
				{
					KeyTag:     10,
					Algorithm:  3,
					DigestType: 5,
					Digest:     "FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102FFAB0102",
				},
			},
		},
	}

	// Generate the response with the default result data and two extensions.
	response := types.Response{
		Result: []types.Result{
			{
				Code:    epp.EppOk.Code(),
				Message: epp.EppOk.Message(),
			},
		},
		ResultData: diResponse,
		// Inline construct an extension type that holds both DNSSEC and IIS.
		Extension: struct {
			types.IISExtensionInfoDataType
			types.DNSSECExtensionInfoDataType
		}{
			diIISExtensionResponse,
			diDNSSECExtensionResponse,
		},
		TransactionID: types.TransactionID{
			ServerTransactionID: "ABC-123",
		},
	}

	return epp.Encode(
		response,
		epp.ServerXMLAttributes(),
	)
}

func createDomain(s *epp.Session, data []byte) ([]byte, error) {
	dc := struct {
		Data types.DomainCreate `xml:"command>create>create"`
	}{}

	if err := xml.Unmarshal(data, &dc); err != nil {
		return nil, err
	}

	// Do stuff with dc which holds all (validated) domain create data.
	spew.Dump(dc)

	return epp.Encode(
		epp.CreateErrorResponse(epp.EppUnimplementedCommand, "not yet implemented"),
		epp.ServerXMLAttributes(),
	)
}

func createContactWithExtension(s *epp.Session, data []byte) ([]byte, error) {
	cc := struct {
		types.ContactCreate
		types.IISExtensionCreate
	}{}

	if err := xml.Unmarshal(data, &cc); err != nil {
		return nil, err
	}

	// Do stuff with cc which holds all (validated) domain create data.

	return epp.Encode(
		epp.CreateErrorResponse(epp.EppUnimplementedCommand, "not yet implemented"),
		epp.ServerXMLAttributes(),
	)
}

func verifyClientCertificate(certs []*x509.Certificate) error {
	if len(certs) != 1 {
		return errors.New("dind't find one single client ceritficate")
	}

	cert := certs[0]
	_, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("could not convert public key")
	}

	// Do something with public key.
	return nil
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
