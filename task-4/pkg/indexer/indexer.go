package indexer

import (
	"regexp"
	"strings"
	"task-4/pkg/spider"
)

func CreateIndex(data map[string]spider.PageData) map[string][]string {
	index := make(map[string][]string)

	for url, pageData := range data {
		pageWords := words(pageData.Text)

		for i := range pageWords {
			word := pageWords[i]
			if containString(index[word], url) {
				continue
			}

			index[word] = append(index[word], url)
		}
	}

	return index
}

func words(text string) []string {
	wordsSeparatorRegexp := regexp.MustCompile("[\\w]+")
	return wordsSeparatorRegexp.FindAllString(text, -1)
}

func containString(words []string, word string) bool {
	for i := range words {
		if strings.Compare(words[i], word) == 0 {
			return true
		}
	}

	return false
}
