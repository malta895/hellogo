package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetchedUrlsCache sync.Map

type result struct {
	body string
	urls []string
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
	var body string
	var urls []string
	if cachedResult, ok := fetchedUrlsCache.Load(url); ok {
		castedResult := cachedResult.(result)
		body = castedResult.body
		urls = castedResult.urls
	} else {
		var err error
		body, urls, err = fetcher.Fetch(url)
		if err != nil {
			return fmt.Sprintln(err)
		}
		fetchedUrlsCache.Store(url, result{
			body, urls,
		})
	}

	foundUrl := fmt.Sprintf("found: %s %q\n", url, body)
	alreadyFetched := foundUrl
	for _, u := range urls {
		alreadyFetched += Crawl(u, depth-1, fetcher)
	}
	return alreadyFetched
}
