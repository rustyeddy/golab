package walker

import (
	"fmt"
	"log"
	"strings"
	"golang.org/x/net/html"
)

// Crawl a URL and return list set of links (no we need the dom)
func (s *Site) Crawl(url string) []string {

	fmt.Printf("Crawling %s\n", url)

	// Do not crawl the page if we already have
	if _, ok := s.Pages[url]; ok {
		return nil
	}
	page := NewPage(s, url)

	if s.Pages == nil {
		s.Pages = make(map[string]Page)
	}

	slen := len(s.BaseUrl)
	if !strings.EqualFold(url[:slen], s.BaseUrl) {
		fmt.Printf("Page NO match: %s - %s\n", url[:slen], s.BaseUrl)
		if _, ok := page.External[url]; ok {
			page.External[url]++
		} else {
			//page.External := make(map[string]int)
			//page.External[url] = 1
		}

	}
	s.Pages[url] = *page
	links, err := page.ExtractLinks(url)
	if err != nil {
		log.Print(err)
	}
	return links

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// TODO: replace this with an anonymous function and closure
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
