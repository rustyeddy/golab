package domains

import (
	"fmt"
	"strconv"
)

// Domain structure represents a single domain name as leased
// from some Registrar (Namecheap in our case).
type Domain struct {
	ID        int64  `csv:"ID"`       // ID unique to the provider (Registrar)
	Name      string `csv:Name`       // The domain name
	Registrar string `csv:Regsitrar`  // Regsitrar that we bought domain from
	Created   string `csv:Created`    // Date the domain was Created
	Expires   string `csv:Expires`    // Date the domain expires
	Expired   bool   `csv:Expired`    // is this domain expired? (from registrar)
	Locked    bool   `csv:Locked`     // is this domain locked
	Autorenew bool   `csv:Autorenew ` // will it auto renew before it expires?
}

// NewDomain creates a new dowmain from an array of strings
func NewDomain(data []string, headers []string) (d *Domain) {
	id, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Errorf("failed to create a new domain for %s", data[0])
		return nil
	}

	// Convert strings to Domain field types
	ex, e1 := strconv.ParseBool(data[4])
	lkd, e2 := strconv.ParseBool(data[5])
	ren, e3 := strconv.ParseBool(data[6])
	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Printf("failed to parse bool for one of %s - %s - %s", data[4], data[5], data[6])
		return nil
	}
	d = &Domain{
		ID:        int64(id),
		Name:      data[1],
		Created:   data[2],
		Expires:   data[3],
		Expired:   ex,
		Locked:    lkd,
		Autorenew: ren,
	}
	return d
}

// DomainDB provides thread-safe access to a database of domains
type DomainDB interface {

	// ListDomains returns a table of domains (index of map?)
	ListDomains() (*DomainTable, error)

	// GetDomain get a specific domain
	GetDomain(name string) (*Domain, error)

	// AddDomain saves a domain in our database
	AddDomain(d *Domain) (id int64, err error)

	// DelDomain removes the domain from the database
	DelDomain(name string) error

	// UpdateDomain changes information in an existing domain entity
	UpdateDomain(d *Domain) error

	// Close closes the database
	Close()
}
