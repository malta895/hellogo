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

type fetchOkResult struct {
	body string
	urls []string
	err  error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) string {
	// TODO: Fetch URLs in parallel.
	if depth <= 0 {
		return ""
	}
	var body string
	var urls []string
	if cachedResult, ok := fetchedUrlsCache.Load(url); ok {
		castedResult := cachedResult.(fetchOkResult)
		if castedResult.err != nil {
			return fmt.Sprintln(castedResult.err)
		}
		body = castedResult.body
		urls = castedResult.urls
	} else {
		var err error
		body, urls, err = fetcher.Fetch(url)
		if err != nil {
			fetchedUrlsCache.Store(url, fetchOkResult{
				"", nil, err,
			})
			return fmt.Sprintln(err)
		}
		fetchedUrlsCache.Store(url, fetchOkResult{
			body, urls, nil,
		})
	}

	foundUrl := fmt.Sprintf("found: %s %q\n", url, body)
	alreadyFetched := foundUrl
	for _, u := range urls {
		alreadyFetched += Crawl(u, depth-1, fetcher)
	}
	return alreadyFetched
}
