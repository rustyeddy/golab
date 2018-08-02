package main

import (
	"time"
	"net"
	"os"
	"io"
	log "github.com/rustyeddy/logrus"
)

func TimeClient(hostport string) {
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		log.Fatalf(" %s, %v", hostport, err)
	}
	defer conn.Close()
	done := make(chan bool)
	go func(c net.Conn) {
		io.Copy(os.Stdout, conn) // XXX Ignoring errors
		log.Println("done")
	}(conn)
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Fatalln(err)
	}
	done <- true
}

func hndlTimeConn(c net.Conn) {
	defer c.Close()
	for {
		t := time.Now()
		_, err := io.WriteString(c, t.Format(time.RFC3339) + "\n")
		if err != nil {
			log.Errorln("hndlTimeConn", err)
		}
		time.Sleep(1 * time.Second)
	}
}
