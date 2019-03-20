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
	mux.AddHandler("command/create/domain", createDomain)

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
				Version: []string{"1.0"},
			},
		},
	}

	return epp.Encode(greeting, epp.ServerXMLAttributes(), "")
}

func login(s *epp.Session, data []byte) ([]byte, error) {
	login := types.Login{}
	if err := xml.Unmarshal(data, &login); err != nil {
		return nil, err
	}

	// Authenticate the user found in login type.

	return epp.Encode(
		epp.CreateErrorResponse(
			epp.EppOk,
			fmt.Sprintf("user '%s' signed in with password '%s'", login.ClientID, login.Password),
		),
		epp.ServerXMLAttributes(),
		"",
	)
}

func createDomain(s *epp.Session, data []byte) ([]byte, error) {
	dc := types.DomainCreate{}
	if err := xml.Unmarshal(data, &dc); err != nil {
		return nil, err
	}

	// Do stuff with dc which holds all (validated) domain create data.

	return epp.Encode(
		epp.CreateErrorResponse(epp.EppUnimplementedCommand, "not yet implemented"),
		epp.ServerXMLAttributes(),
		"",
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
