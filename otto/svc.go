package main

import (
	"net"
	log "github.com/rustyeddy/logrus"
)

type Service struct {
	name string
	hostport string
	enabled bool
}

func StartService(name, hostport string, done chan<- bool, cb func(c net.Conn)) {
	ln, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatalf("tcp %s err ", hostport, err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("listen %s %s %v", cfg.Proto, cfg.Hostport)
			continue
		}
		go cb(conn)
	}
	done <- true
}
