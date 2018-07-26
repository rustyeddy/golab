package golib

import (
	"log"
	"net"
	"os"
	"time"
)

/*
## Structs

- DNSHostRecord contains DNS records for a specific host
  - Hostname string, IP, CNAME, TXT, MX, NS, ...

## Exported API

- DNSHostRecord(host string) *DNSHostRecord
*/

// DNSHostRecord to hold all known DNS records
type DNSHostRecord struct {
	Host      string
	Addrs     []string
	IP        []net.IP
	CNAME     string
	TXT       []string
	MX        []*net.MX
	NS        []*net.NS
	Refreshed time.Time
	Err       error
}

func dnsit() {
	// make sure the caller provided us with at least a
	// command to run
	if len(os.Args) < 3 {
		log.Printf("give me a command.  I don't know what to do.. ")
		os.Exit(1)
	}

	// Get our command, we currently only support / understand cmd ...
	cmd, host := os.Args[1], os.Args[2]
	if cmd == "" || host == "" {
		log.Printf("something fishy with my args expected non empty strings got:")
		log.Printf("\tArg[1] (cmd) -> %s", cmd)
		log.Printf("\tArg[2] (host) -> %s", host)
		os.Exit(3)
	}

	// Debug message possibility right here
	switch cmd {
	case "host":
		recs := DNSHostRecords(host)
		if recs == nil {
			log.Fatalf("failed to get host records")
		}
		log.Printf("%v", recs)

	default:
		log.Printf("unknown command %s", cmd)
		os.Exit(3)
	}
}

// DNSHostRecords will collect all DNSHost records: IP, CNAME, Host, NS, MX, TXT
func DNSHostRecords(host string) (d *DNSHostRecord) {

	d = new(DNSHostRecord)
	d.Host = host
	ip, err := net.LookupIP(host)
	if err == nil {
		d.IP = ip
	}
	cname, err := net.LookupCNAME(host)
	if err == nil {
		d.CNAME = cname
	}
	addrs, err := net.LookupHost(host)
	if err == nil {
		d.Addrs = addrs
	}
	mx, err := net.LookupMX(host)
	if err == nil {
		d.MX = mx
	}
	ns, err := net.LookupNS(host)
	if err == nil {
		d.NS = ns
	}
	txt, err := net.LookupTXT(host)
	if err == nil {
		d.TXT = txt
	}
	return d
}
