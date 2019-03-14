package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	epp "github.com/bombsimon/epp-go"
)

func main() {
	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{generateCertificate()},
		},
	}

	greeting, err := client.Connect(":4701")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(greeting))

	fmt.Println("> Automatic login!")
	time.Sleep(1 * time.Second)

	response, err := client.Login("some-user", "some-password")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(response))

	for {
		scnr := bufio.NewScanner(os.Stdin)
		scnr.Split(func(data []byte, _ bool) (int, []byte, error) {
			if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
				return i + 2, data[0:i], nil
			}

			return 0, nil, nil
		})

		scnr.Scan()
		data := scnr.Bytes()
		if scnr.Err() != nil {
			log.Fatal(scnr.Err().Error())
		}

		if string(data) == "exit" {
			return
		}

		response, err := client.Send(data)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(string(response))
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
