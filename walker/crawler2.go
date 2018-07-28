package walker

import (
	"flag"
	"fmt"
)

var (
	// Read command line args
	url     = flag.String("url", "http://rustyeddy.com", "url http://example.com")
	Version string
	Build   string
)

func main() {
	help()

	flag.Parse()
	fmt.Printf("Accessing site %s\n", *url)

	site := Site{BaseUrl: *url}
	urls := []string{site.BaseUrl}

	page := Page{Url: "http://RustyEddy.com"}
	page.BreadthFirst(site.Crawl, urls)
}

func help() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)
}
