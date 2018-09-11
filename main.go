package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var (
	root = "http://172.16.33.11/Person/"
	maxDepth = 1
	selector = "a[href$=dcm]"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// MaxDepth is 1, so only the links on the scraped page
		// is visited, and no further links are followed
		colly.MaxDepth(maxDepth),
		colly.Async(true),
	)


	// On every a element which has href attribute call callback
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		e.Request.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(root)
	c.Wait()
}