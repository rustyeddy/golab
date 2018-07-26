package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/mmcdole/gofeed"
)

type Parameters struct {
	Url         string
	Query       string
	FullRequest string
}

// Args are a specific invocation of parameters
var Args Parameters

func init() {
	flag.StringVar(&Args.Url, "url", "https://www.upwork.com/jobs/atom", "job listing url")
	flag.StringVar(&Args.Query, "q", "", "query to append to url")
}

func main() {

	// will print Usage and exit if it fails
	flag.Parse()

	// Vett the provided URL returning the request url or error
	requrl := GetURL(Args.Url, Args.Query)
	if requrl == "" || Args.FullRequest != requrl {

		// Hmmm. I quickly hit a query limit.  Need to figure out what
		// that is.
		log.Println("failed to build a url correctly ... ")
		log.Printf("\n\turl  %s\n\tq    %s\n\tfull %s", Args.Url, Args.Query, requrl)
	}

	// Let's see if we can get some jobs!
	log.Printf("Getting some work... ")
	feed, err := GetFeed(requrl)
	if feed == nil {
		log.Fatal(err)
	}

	// Just print to console .. or log for now
	PrintFeed(feed)
}

// GetFeed returns a Feed which contains (amoung) other
// things FeedItems .. We will walk the FeedItems and print briefly
func GetFeed(requrl string) (f *gofeed.Feed, err error) {
	fp := gofeed.NewParser()
	if fp == nil {
		return nil, fmt.Errorf("failed to get an atom parser. bye...")
	}
	return fp.ParseURL(requrl)
}

// PrintFeed will print a summary of the feed we just snarfed
func PrintFeed(f *gofeed.Feed) {
	fmt.Println(f.Title)
	fmt.Printf("updated %s -- published %s\n", f.Updated, f.Published)
	fmt.Println("----------------------------\n")
	fmt.Println("desc: ", f.Description)
	fmt.Println("----------------------------\n")
	fmt.Println("link: ", f.Link)
	fmt.Println("categories: ")
	for _, s := range f.Categories {
		fmt.Println("\t" + s)
	}

	fmt.Println(" = =============== ITEMS ================ = ")
	// Now start printing the actual items
	for _, it := range f.Items {
		fmt.Printf("%v", it)
	}
}

// Get a URL that has been vetted and combined based on two possible
// strings for baseurl and query.  Request URL is constructed as:
/*

% work -url http://www.example.com/  -q <qstr>
    requrl => http://www.example.com/atom?q=python&api=true

% work -url http://www.example.com/atom?q=<qstr>
    requrl => http://www.example.com/atom?q=python

% work -url http://www.example.com/auom
    requrl => http://www.example.com/auom  !!! Error no query

if you supply a -url it will only be modified if you have also supplied
a -q for the query.  NOTE: The -q flag does NOT have a default.

There fore there are three combinations of invocations for this function
that are legal:

  url -: url must include query string in URL (correctly formatted)
  url q: url must NOT have a query string (or it will be replaced by -q)
   -  q: url will append qstr efault baseurl (no query)
   -  -: url will use baseurl, however we have no query, no default fails

A valid, legal, safe URL will be returned or nothing (in event of a problem).
*/

func GetURL(base, query string) string {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatal(err)
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	if u.Host == "" {
		// make this a config item
		u.Host = "www.upwork.com"
	}
	q := u.Query()
	if query != "" {
		q.Set("q", query)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
