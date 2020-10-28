package main

import (
	"fmt"
	"task-4/pkg/indexer"
	"task-4/pkg/spider"
)

type Scanner interface {
	Scan() (map[string]spider.PageData, error)
}

func main() {
	s := spider.Spider{
		Url:   "https://golang.org",
		Depth: 3,
	}

	parseSite(s)
}

func parseSite(s Scanner) {
	var searchKeyword string

	pagesData, err := s.Scan()

	if err != nil {
		return
	}

	fmt.Printf("Scan completed, %v sites found\n", len(pagesData))

	index := indexer.CreateIndex(pagesData)

	for {
		fmt.Print("Please enter search keyword (exit - ^C): ")
		fmt.Scanf("%s", &searchKeyword)

		urls, ok := index[searchKeyword]

		if !ok {
			fmt.Println("Nothing found")
		}

		for i := range urls {
			url := urls[i]
			fmt.Printf("Found %v - %v\n", url, pagesData[url].Title)
		}
	}
}
