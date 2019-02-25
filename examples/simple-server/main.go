package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/pkg/errors"

	eppserver "github.com/bombsimon/epp-server"
)

func main() {
	// Create temporary certificates and remove when shutting down.
	// NOTE: This is for testing only and should not be used.
	tempCert, tempKey := createCertificateFiles()
	defer func() {
		os.Remove(tempCert)
		os.Remove(tempKey)
	}()

	server := eppserver.Server{
		IdleTimeout:      5 * time.Minute,
		MaxSessionLength: 10 * time.Minute,
		Addr:             ":4701",
		Handler:          RouteMessage,
		Greeting:         Greet,
	}

	// Support graceful shutdown.
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		server.Stop()
	}()

	log.Println("Running server...")
	if err := server.ListenAndServe(tempCert, tempKey); err != nil {
		log.Fatal(err.Error())
	}
}

func Greet(s *eppserver.Session) ([]byte, error) {
	return []byte("greetings!"), nil
}

func RouteMessage(s *eppserver.Session, data []byte) ([]byte, error) {
	xmlData, err := xmlquery.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	xpaths := []string{"/epp/hello",
		"/epp/command/login",
		"/epp/command/logout",
		"/epp/command/poll",
		"/epp/command/check",
		"/epp/command/create",
		"/epp/command/delete",
		"/epp/command/info",
		"/epp/command/renew",
		"/epp/command/transfer",
		"/epp/command/update",
	}

	for _, xpath := range xpaths {
		if xpath == "/epp/command/login" {
			err := verifyClientCertificate(s.ConnectionState().PeerCertificates)
			if err != nil {
				return []byte("could not verify peer certificates"), nil
			}
		}
		node := xmlquery.FindOne(xmlData, xpath)
		if node != nil {
			return []byte(fmt.Sprintf("got message %s", xpath)), nil
		}
	}

	return []byte("no handler for this"), nil
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

func createPrivKey(priv *rsa.PrivateKey) string {
	f, _ := ioutil.TempFile(".", "priv*.key")
	defer f.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}

	_ = pem.Encode(f, privateKey)

	return f.Name()
}

func createCertificate(priv *rsa.PrivateKey, pub rsa.PublicKey) string {
	f, _ := ioutil.TempFile(".", "pub*.pem")
	defer f.Close()

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

	certificate, _ := x509.CreateCertificate(rand.Reader, cert, cert, &pub, priv)

	certFile := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate,
	}

	_ = pem.Encode(f, certFile)

	return f.Name()
}

func createCertificateFiles() (string, string) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	tempKey := createPrivKey(key)
	tempCert := createCertificate(key, key.PublicKey)

	return tempCert, tempKey
}
