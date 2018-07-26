package providers

import (
	"fmt"
	"io/ioutil"
	"log"

	nc "github.com/rustyeddy/go-namecheap"
	yml "gopkg.in/yaml.v2"
)

/*
	// DomainProvider defines the expectations from a Registrar
	type DomainProvider interface {
		GetDomains() *Domains
		GetDomain(name string)

		GetNameservers(domain string) []string
		SetNameservers(domain string, ns []string) error

		CheckDomain(domain string) bool
		PurchaseDomain(domain string) error
		Renew(domain string) bool
	}

	// DNSProvider manages DNS records or us
	type DNSProvider interface {
		GetDNSRecords() ([]DNSRecord, error)
		GetDNSRecord(dname, rect string) (error, []DNSRecord)
		SetDNSRecord(dname, record, data string) error
	}

	type DNSRecord struct {
		RecType  string
		RecKey   string
		RecValue string
	}
*/

// Namecheap - A Registrar
type Namecheap struct {
	name   string
	client *nc.Client
	err    error
}

// Credentials for namecheap API
type credentials struct {
	APIUser  string `yaml:"APIUser"`
	APIToken string `yaml:"APIToken"`
}

// NewNamecheap create a new Namecheap client
func NewNamecheap(credspath string) (cli Namecheap) {
	cli = Namecheap{name: "Namecheap"}

	creds, err := getCreds(credspath)
	if err != nil {
		log.Fatal("expected creds but got an error %v", err)
	}

	// connect with namecheap
	cli.client = nc.NewClient(creds.APIUser, creds.APIToken, creds.APIUser)
	if cli.client == nil {
		log.Printf("yikes namecheap client failed")
		cli.err = fmt.Errorf("client connect returned nil")
	}
	return cli
}

// getCreds from config to access API. Used only be connect.
func getCreds(credpath string) (c credentials, err error) {
	var data []byte

	// Now read creds from credpath file.
	if data, err = ioutil.ReadFile(credpath); err != nil {
		log.Printf("failed to read file %s with %v\n", credpath, err)
		return c, err
	}

	// unravel the yaml to get creds
	if err := yml.Unmarshal(data, &c); err != nil {
		log.Printf("reading yml %s failed %v\n", credpath, err)
		return c, err
	}
	return c, err
}

// Name of the registrar
func (n *Namecheap) Name() string {
	return n.name
}

// GetDomains returns the domains managed by this registrar.  We will
// check for a cached version of the domain list before fetching from
// Namecheap
func (n *Namecheap) GetDomains() []nc.DomainGetListResult {

	// get the domList from namecheap
	dlist, err := n.client.DomainsGetList()
	if err != nil {
		log.Printf("failed to get domains %v", err)
		return nil
	}
	return dlist
}
