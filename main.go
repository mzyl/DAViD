package main

import (
  "fmt"

  "github.com/gocolly/colly"
)

func main() {
  // instantiate default collector 
  c := colly.NewCollector(
    // visit only domains: hackerspaces.org, wiki.hackerspaces.org
    colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
  )

  // on every a element which has href attribute call callback
  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    link := e.Attr("href")
    // print link
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    // visit link found on page
    // only those links are visited which are in AllowedDomains
    c.Visit(e.Request.AbsoluteURL(link))
  })

  // before making a request print "Visiting..."
  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL.String())
  })

  // start scraping on https://hackerspaces.org
  c.Visit("https://hackerspaces.org/")
}
