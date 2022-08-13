package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) string {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return ""
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		return fmt.Sprintln(err)
	}
	foundUrl := fmt.Sprintf("found: %s %q\n", url, body)
	alreadyFetched := foundUrl
	for _, u := range urls {
		alreadyFetched += Crawl(u, depth-1, fetcher)
	}
	return alreadyFetched
}
