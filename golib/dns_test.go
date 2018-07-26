package golib

import (
	"log"
	"testing"
	//log "github.com/rustyeddy/logrus"
)

func TestGetDNSHost(t *testing.T) {
	host := "usc.edu"
	dm := DNSHostRecords(host)
	if dm == nil {
		t.Fatalf("expected a DNSHostRecord got none")
	}
	if dm.IP == nil {
		t.Error("expected an IP address got none")
	}
	if dm.CNAME == "" {
		t.Errorf("expected a CNAME for %s", host)
	}
}

func TestDNSHostRecords(t *testing.T) {
	t.Skip()
	type args struct {
		Host string
	}
	tests := []struct {
		name  string
		args  args
		wantD *DNSHostRecord
	}{
		{"usc", args{"usc.edu"}, &DNSHostRecord{Host: "usc.edu"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD := DNSHostRecords(tt.args.Host)
			log.Println("~~~~~ DNS records for USC ~~~~")
			log.Printf("\t%+v", gotD)
			log.Println("~~~~~~    The End ~~~~~~~~~~~~")
		})
	}
}
