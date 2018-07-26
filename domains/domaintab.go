package domains

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rustyeddy/domains/providers"
)

// DomainTable Domains indexed by name
type DomainMap map[string]Domain

func NewDomainMap(c int) DomainMap {
	return make(DomainMap, c)
}

// in mem table of Domains
type DomainTable struct {
	Name string
	DomainMap
}

// getDomains() will first check for the local csv file with domains.
// If that does not exist, it will go to the provider(s)
func GetDomains() (dt *DomainTable, err error) {

	exitOnFailure := true

	// Try database first
	dt, err = GetDomainsDB()
	if err != nil {
		err = fmt.Errorf("expected domains but got error %v", err)
		if exitOnFailure {
			return nil, err
		}
		log.Println(err)
	}

	if _, err = os.Stat(DomainsPath); err != nil {
		log.Println(fmt.Errorf("expected to stat %s but got an error %v", DomainsPath, err))
	}
	if dt, err = GetDomainsProvider(CredsFile); err != nil {
		return nil, fmt.Errorf("expected get domains db but got an error %v", DomainsPath, err)
	}
	return dt, err
}

// getDomainsLocal retrieves the domains from a local cache, if
// one exists
func GetDomainsCSV(lpath *string) (dt *DomainTable, err error) {
	csvrecs, err := ReadCSVFile(lpath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file %v", err)
	}
	dt = GetDomainsRecs(csvrecs)
	if dt == nil {
		return nil, fmt.Errorf("failed to get domains from CSV")
	}
	return dt, err
}

// GetDomainsCSV loads domains from a CSV structure (previously
// read from a csv file)
func GetDomainsRecs(csvrecs [][]string) (dt *DomainTable) {
	flds := csvrecs[0]
	dt = &DomainTable{
		Name:      "csv",
		DomainMap: make(DomainMap, len(csvrecs)),
	}
	dm := dt.DomainMap
	for _, record := range csvrecs[1:] {
		dom := NewDomain(record, flds)
		dm[dom.Name] = *dom
	}
	return dt
}

// GetDomains from namecheap
func GetDomainsProvider(creds string) (domains *DomainTable, err error) {
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	nc := providers.NewNamecheap(creds)

	// no local variable, go get it from NC
	if ncdomains := nc.GetDomains(); ncdomains != nil {
		nclen := len(ncdomains)
		domains = &DomainTable{
			Name:      "namecheap",
			DomainMap: make(DomainMap, nclen+1),
		}
		header := []string{
			"ID", "Name", "Created", "Expires", "Expired", "Locked", "Autorenew",
		}

		// write the header out
		if err := w.Write(header); err != nil {
			log.Fatalln("failed to write csv header ", err)
		}

		// Convert to domains from namecheap
		for _, ncdom := range ncdomains {
			// index the domain in the DomainTable by name
			d := Domain{
				ID:        int64(ncdom.ID),
				Name:      ncdom.Name,
				Created:   ncdom.Created,
				Expires:   ncdom.Expires,
				Expired:   ncdom.IsExpired,
				Locked:    ncdom.IsLocked,
				Autorenew: ncdom.AutoRenew,
			}
			domains.DomainMap[ncdom.Name] = d

			csv := []string{
				strconv.FormatInt(int64(ncdom.ID), 10),
				ncdom.Name,
				ncdom.Created,
				ncdom.Expires,
				strconv.FormatBool(ncdom.IsExpired),
				strconv.FormatBool(ncdom.IsLocked),
				strconv.FormatBool(ncdom.AutoRenew),
			}

			// write out the records now
			if err := w.Write(csv); err != nil {
				log.Printf("ERROR -> writing csv file for domain %s", ncdom.Name)
				log.Printf("ERROR ->         error %v", err)
			}
		}
	}
	return domains, err
}
