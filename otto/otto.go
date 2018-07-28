package main

import (
	"fmt"
	"time"
	"flag"
	"net"
	"io"
	log "github.com/rustyeddy/logrus" 
)

type Configuration struct {
	Hostport string // "<hostname>:<port>"
	Proto string
}

var (
	cfg Configuration
)

func init() {
	flag.StringVar(&cfg.Hostport, "h", "localhost:1231", "set the server host")
	flag.StringVar(&cfg.Proto, "P", "tcp", "TCP, HDP or ...")
}

func main() {
	fmt.Println("Starting spinner... ")

	listener, err := net.Listen("tcp", "localhost:1231")
	if err != nil {
		log.Fatalf("tcp localhost:1231 err ", err)
	}

	for {
		conn, err := net.Listen(cfg.Proto, cfg.Hostport)
		if err != nil {
			log.Fatalf("listen %s %s %v", cfg.Proto, cfg.Hostport)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("14:04:04\n"))
		if err != nil {
			return 
		}
		time.Sleep(1 * time.Second)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// fib
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

