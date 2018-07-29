package main

import (
	"fmt"
	"flag"
	log "github.com/rustyeddy/logrus" 
)

// Context has all the variables we need to satisfy a request
type Context struct {
	Configuration
}

// Configuration items are variables that can be set in a configuration
// file at startup time (or read in during run time)
type Configuration struct {
	Hostport string // "<hostname>:<port>"
	Proto string // tcp, udp, http ..
	TimeServer bool
}

var (
	cfg Configuration
)

func init() {
	flag.StringVar(&cfg.Hostport, "h", "localhost:1231", "set the server host")
	flag.StringVar(&cfg.Proto, "P", "tcp", "TCP, HDP or ...")
	flag.BoolVar(&cfg.TimeServer, "time", true, "start the time server port 1231 by default ")

	log.Debug("Init complete")
}

func main() {
	fmt.Println("OttO ... ")
	StartTimeService()
}


