package main

import (
	"time"
	"net"
	"io"
	log "github.com/rustyeddy/logrus"
)

func StartTimeService() (err error) {

	ln, err := net.Listen("tcp", "localhost:1231")
	if err != nil {
		log.Fatalf("tcp localhost:1231 err ", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("listen %s %s %v", cfg.Proto, cfg.Hostport)
			continue
		}
		go handleTimeSvcConn(conn)
	}
	
	return nil
}

func handleTimeSvcConn(c net.Conn) (err error) {
	defer c.Close()
	for {
		t := time.Now()
		_, err := io.WriteString(c, t.Format(time.RFC3339) + "\n")
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}
