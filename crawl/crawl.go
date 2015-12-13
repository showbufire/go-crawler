package crawl

import (
	"sync"

	"github.com/showbufire/crawler/links"
)

const (
	maxWorkers = 5
)

type job struct {
	url   string
	depth int
}

type jobResult struct {
	urls  []string
	depth int
	err   error
}

func Crawl(url string, maxdepth int) []string {
	res := make(chan *jobResult)
	jobs := make(chan *job)
	wg := sync.WaitGroup{}

	for i := 0; i < maxWorkers; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				ls, err := links.Extract(j.url)
				if err != nil {
					res <- &jobResult{err: err}
				} else {
					res <- &jobResult{urls: ls, depth: j.depth - 1}
				}
			}
		}()
	}
	go func() {
		res <- &jobResult{
			urls:  []string{url},
			depth: maxdepth,
		}
	}()
	visited := make(map[string]int)
	for n := 1; n > 0; n -= 1 {
		r := <-res
		if r.err != nil {
			continue
		}
		for _, u := range r.urls {
			if _, ok := visited[u]; !ok {
				visited[u] = r.depth
				if r.depth > 0 {
					n += 1
					go func() {
						jobs <- &job{url: u, depth: r.depth}
					}()
				}
			}
		}
	}
	close(jobs)
	defer wg.Wait()

	urls := []string{}
	for k, _ := range visited {
		urls = append(urls, k)
	}
	return urls
}
