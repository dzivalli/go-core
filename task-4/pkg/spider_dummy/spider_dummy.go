package spider_dummy

import "task-4/pkg/spider"

type Spider struct {
	Url   string
	Depth int
}

func (s Spider) Scan() (data map[string]spider.PageData, err error) {
	data = make(map[string]spider.PageData)

	data["https://golang.org"] = spider.PageData{
		Title: "Golang title",
		Text:  "Golang is super cool",
	}

	data["https://go.dev"] = spider.PageData{
		Title: "Go Dev title",
		Text:  "Go Dev text",
	}

	return data, nil
}
