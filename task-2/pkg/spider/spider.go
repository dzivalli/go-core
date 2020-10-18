package spider

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type pageData struct {
	Title string
	Text  string
}

func Scan(url string, depth int) (data map[string]pageData, err error) {
	data = make(map[string]pageData)

	parse(url, depth, data)

	return data, nil
}

func parse(url string, depth int, data map[string]pageData) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}
	text := pageText(page, []string{})

	data[url] = pageData{
		Title: pageTitle(page),
		Text:  strings.Join(text, ""),
	}

	links := pageLinks(nil, page)
	for _, link := range links {
		_, exists := data[link]

		if !exists && strings.HasPrefix(link, "http") {
			parse(link, depth-1, data)
		}
	}

	return nil
}

func pageText(n *html.Node, words []string) []string {
	if n.Type == html.TextNode {
		words = append(words, n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		return words
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words = pageText(c, words)
	}

	return words
}

func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild == nil {
			return "Untitled"
		} else {
			return n.FirstChild.Data
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
