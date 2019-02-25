package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	eppserver "github.com/bombsimon/epp-server"
)

func main() {
	conn, err := tls.Dial("tcp", ":4701", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		msg, err := eppserver.ReadMessage(conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			log.Fatal(err.Error())
		}
		fmt.Println(string(msg))

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

		err = eppserver.WriteMessage(conn, data)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
