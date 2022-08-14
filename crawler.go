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

type fetchResult struct {
	body string
	urls []string
	err  error
}

var fetchedUrlsCache sync.Map

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) string {
	var crawlHelper func(url string, depth int, fetcher Fetcher, fetchedChannel chan string, quit chan int)

	crawlHelper = func(url string, depth int, fetcher Fetcher, fetchedChannel chan string, quit chan int) {
		if depth <= 0 {
			quit <- 0
			return
		}

		var body string
		var urls []string

		if cachedResult, ok := fetchedUrlsCache.Load(url); ok {
			castedResult := cachedResult.(fetchResult)
			if castedResult.err != nil {
				fetchedChannel <- fmt.Sprintln(castedResult.err)
				return
			}
			body = castedResult.body
			urls = castedResult.urls
		} else {
			var err error
			body, urls, err = fetcher.Fetch(url)
			if err != nil {
				fetchedUrlsCache.Store(url, fetchResult{
					"", nil, err,
				})
				fetchedChannel <- fmt.Sprintln(err)
				return
			}
			fetchedUrlsCache.Store(url, fetchResult{
				body, urls, nil,
			})
		}
		foundUrl := fmt.Sprintf("found: %s %q\n", url, body)
		fetchedChannel <- foundUrl
		for _, u := range urls {
			go crawlHelper(u, depth-1, fetcher, fetchedChannel, quit)
		}
	}

	fetchedChannel, quit := make(chan string), make(chan int)
	var foundUrls string
	go crawlHelper(url, depth, fetcher, fetchedChannel, quit)
	for {
		select {
		case foundUrl := <-fetchedChannel:
			fmt.Printf("Found url %v\n", foundUrl)
			foundUrls += foundUrl
		case <-quit:
			fmt.Println("Received quit signal")
			return foundUrls
		default:
			fmt.Println("default case of select")
		}
	}
	return foundUrls
}
