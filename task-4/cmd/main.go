package main

import (
	"fmt"
	"strings"
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
	var found bool

	sites, err := s.Scan()

	if err != nil {
		return
	}

	fmt.Printf("Scan completed, %v sites found\n", len(sites))

	for {
		found = false
		fmt.Print("Please enter search keyword (exit - ^C): ")
		fmt.Scanf("%s", &searchKeyword)

		for url, pageData := range sites {
			if strings.Index(pageData.Text, searchKeyword) > 0 {
				found = true
				fmt.Printf("Found %v - %v\n", url, pageData.Title)
			}
		}

		if !found {
			fmt.Println("Nothing found")
		}
	}
}
