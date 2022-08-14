package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var fakeFetcherCalls = make(map[string]int)

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if _, ok := fakeFetcherCalls[url]; ok {
		fakeFetcherCalls[url]++
	} else {
		fakeFetcherCalls[url] = 1
	}
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func TestCorrectUrlsCrawledAndCached(t *testing.T) {
	found := Crawl("https://golang.org/", 4, fetcher)
	expected := `found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
`
	assert.Equal(t, expected, found)

	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/cmd/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/fmt/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/os/"])
}

func TestCorrectUrlsCrawledAndCachedOrderIgnored(t *testing.T) {
	found := Crawl("https://golang.org/", 4, fetcher)
	expectedUrls := []string{
		`found: https://golang.org/ "The Go Programming Language"`,
		`found: https://golang.org/pkg/ "Packages"`,
		`found: https://golang.org/ "The Go Programming Language"`,
		`found: https://golang.org/pkg/ "Packages"`,
		`not found: https://golang.org/cmd/`,
		`not found: https://golang.org/cmd/`,
		`found: https://golang.org/pkg/fmt/ "Package fmt"`,
		`found: https://golang.org/ "The Go Programming Language"`,
		`found: https://golang.org/pkg/ "Packages"`,
		`found: https://golang.org/pkg/os/ "Package os"`,
		`found: https://golang.org/ "The Go Programming Language"`,
		`found: https://golang.org/pkg/ "Packages"`,
		`not found: https://golang.org/cmd/`,
	}
	for _, expectedUrl := range expectedUrls {
		assert.Contains(t, found, expectedUrl)
	}
	assert.NotContains(t, found, `found: https://golang.org/cmd/ ""`)

	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/fmt/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/pkg/os/"])
	assert.Equal(t, 1, fakeFetcherCalls["https://golang.org/cmd/"])

}
