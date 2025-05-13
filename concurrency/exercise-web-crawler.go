package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of the url and a slice of urls found on that page
	Fetch(url string) (body string, urls []string, err error)
}

var visited = struct {
	m map[string]bool
	sync.Mutex
}{m: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// fetch urls in parallel using a wait group
	// increment the wait group when we call go Crawl()
	// decrement from the group once we've executed this func
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// dont fetch the same url twice; if we have visited this url before just continue
	// maps are not thread safe therefore we need to use a mutex to sync access
	visited.Lock()
	if visited.m[url] {
		visited.Unlock()
		return
	}
	visited.m[url] = true
	visited.Unlock()

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}

	return
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

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
