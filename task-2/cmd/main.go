package main

import (
	"fmt"
	"strings"
	"task-2/pkg/spider"
)

func main() {
	var searchKeyword string
	var found bool

	url := "https://golang.org/"
	sites, err := spider.Scan(url, 2)

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
