package main

import (
	"fmt"
	"net"
	"io"
	log "github.com/rustyeddy/logrus"
)

func EchoClient(hostport, msg string) {
	conn, err := net.Dial("tcp", "localhost:1221")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, "\tmsg")
}

func hndlEchoConn(c  net.Conn) {
	if _, err := io.Copy(c, c); err != nil {
		log.Errorf("echo %v", err)
	}
}

