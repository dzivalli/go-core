package spider

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type PageData struct {
	Title string
	Text  string
}

type Spider struct {
	Url   string
	Depth int
}

// parse given url recursively with given depth
func (s Spider) Scan() (data map[string]PageData, err error) {
	data = make(map[string]PageData)

	parse(s.Url, s.Url, s.Depth, data)

	return data, nil
}

func parse(url string, baseurl string, depth int, data map[string]PageData) error {
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

	data[url] = PageData{
		Title: pageTitle(page),
		Text:  strings.Join(text, ""),
	}

	links := pageRelativeLinks(nil, page)
	for _, relativeLink := range links {
		link := baseurl + relativeLink
		if _, ok := data[link]; !ok {
			parse(link, baseurl, depth-1, data)
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

func pageRelativeLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !strings.HasPrefix(a.Val, "http") && !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageRelativeLinks(links, c)
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
