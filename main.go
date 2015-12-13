package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/showbufire/crawler/crawl"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("no argument provided")
	}
	d, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	urls := crawl.Crawl(os.Args[1], d)
	for _, u := range urls {
		fmt.Printf("%v\n", u)
	}
}
