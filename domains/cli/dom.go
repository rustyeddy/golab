package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/rustyeddy/domains"
)

/*
 * Commands:

 dom ls
     => list of domains

*/
var (
	ConfigDir   *string
	CredsFile   *string
	DomainsPath *string
	Debugging   *bool
)

// ================== USAGE ==========================
var Usage = func() {
	fmt.Fprintf(os.Stderr, "usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "only flagsets have PrintDefaults()?\n")
}

// ============== GLOBAL VARIABLES ==================
func init() {
	basestr := "/Users/rusty/.config/domains"
	ConfigDir = flag.String("config", basestr, "configs are stored")
	CredsFile = flag.String("creds", basestr+"/config.yml", "config file")
	DomainsPath = flag.String("domains", basestr+"/domains.csv", "domains.csv")
	Debugging = flag.Bool("debug", false, "turn on debugging")
}

func main() {
	flag.Parse()

	cmd := "ls"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	var (
		domtab, nctab, loctab *DomainTable
		err                   error
	)

	switch cmd {
	case "ls":
		domtab, err = GetDomains()
	case "db":
		domtab, err = GetDomainsDB()
	case "csv":
		loctab, err = GetDomainsCSV(DomainsPath)
	case "nc":
		nctab, err = GetDomainsProvider(*CredsFile)
	case "syncdb":
		if nctab, err = GetDomainsProvider(*CredsFile); err == nil {
			if loctab, err = GetDomainsCSV(DomainsPath); err == nil {
				if err = DBSync(nctab, loctab); err == nil {
					domtab, err = GetDomains()
				}
			}
		}
	}

	// Check for error not that we are past the command. Die without last rites
	if err != nil {
		log.Fatalf("cmd (%s) failed to read domains %v", err)
	}

	if loctab != nil {
		for _, dom := range loctab.DomainMap {
			fmt.Printf("\t%40s - %s\n", dom.Name, dom.Created)
		}
	}
}
