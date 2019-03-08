package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"

	eppserver "github.com/bombsimon/epp-server"
)

func main() {
	mux := eppserver.NewMux()
	mux.AddHandler("command/login", func(s *eppserver.Session, data []byte) ([]byte, error) {
		// Do stuff.
		return []byte("login"), nil
	})
	mux.AddHandler("command/check/contact", func(s *eppserver.Session, data []byte) ([]byte, error) {
		// Do stuff.
		return []byte("contact-check"), nil
	})

	server := eppserver.Server{
		IdleTimeout:      5 * time.Minute,
		MaxSessionLength: 10 * time.Minute,
		Addr:             ":4701",
		Handler:          mux.Handle,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{generateCertificate()},
			ClientAuth:   tls.RequireAnyClientCert,
		},
		Greeting: func(s *eppserver.Session) ([]byte, error) {
			err := verifyClientCertificate(s.ConnectionState().PeerCertificates)
			if err != nil {
				_ = s.Close()

				fmt.Println("could not verify peer certificates")
				return nil, errors.New("could not verify certificates")
			}

			return []byte("greetings!"), nil
		},
	}

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
