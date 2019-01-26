package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	eppserver "github.com/bombsimon/epp-server"
	"github.com/pkg/errors"
)

func main() {
	// Create temporary certificates and remove when shutting down.
	// NOTE: This is for testing only and should not be used.
	tempCert, tempKey := createCertificateFiles()
	defer func() {
		os.Remove(tempCert)
		os.Remove(tempKey)
	}()

	server := eppserver.New("epp.example.test", ":4700", tempKey, tempCert)
	server.VerifyClientCertificateFunc = verifyClientCertificate

	// Support graceful shutdown.
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		server.Stop()
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func verifyClientCertificate(rawCerts [][]byte, _ [][]*x509.Certificate) error {
	if len(rawCerts) != 1 {
		return errors.New("dind't find one single client ceritficate")
	}

	c, err := x509.ParseCertificate(rawCerts[0])
	if err != nil {
		return errors.Wrap(err, "could not parse client certificate")
	}

	_, ok := c.PublicKey.(*rsa.PublicKey)
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
