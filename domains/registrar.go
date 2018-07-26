package domains

// Registrar is the company that we rent our domains from
type Registrar struct {
	Name string
	URL  string
}

// A struct will have to support the calls below to be a registrar
type RegistrarIntf interface {
	GetDomains() *DomainTable

	// TODO - the following calls
	// GetDomain(dname string) *Domain
	// CheckDomain(dname string) bool
	// PurchaseDomain(dname string) bool

	// GetNameservers(dname string) (*Nameservers, error)
	// SetNameservers(hosts *Nameservers) error
}
