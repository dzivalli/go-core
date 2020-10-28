package indexer

import (
	"reflect"
	"strings"
	"task-4/pkg/spider"
	"testing"
)

func Test_words(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Words with space and capital letters",
			args: args{
				text: "Word1 Word2",
			},
			want: []string{"Word1", "Word2"},
		},
		{
			name: "Words with spaces, tabs and new lines",
			args: args{
				text: "   word1    word2 \t  word3 \n  word4    ",
			},
			want: []string{"word1", "word2", "word3", "word4"},
		},
		{
			name: "Words with ",
			args: args{
				text: "   word1.    word2! word3? word4,",
			},
			want: []string{"word1", "word2", "word3", "word4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := words(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("words() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndex(t *testing.T) {
	pagesData := map[string]spider.PageData{
		"https://go.dev": {
			Title: "go dev",
			Text:  "word1 word2 word1",
		},
		"https://go.dev/about": {
			Title: "go about",
			Text:  "word2 word3 word3",
		},
	}

	want := map[string][]string{
		"word1": {
			"https://go.dev",
		},
		"word2": {
			"https://go.dev",
			"https://go.dev/about",
		},
		"word3": {
			"https://go.dev/about",
		},
	}

	got := CreateIndex(pagesData)

	if len(got) != len(want) {
		t.Errorf("Index lengh doesn't match len(got) - %v, len(want) %v", len(got), len(want))
	}

	for word, urls := range want {
		if len(got[word]) != len(urls) {
			t.Errorf("Words don't match got words - %v, want words %v", got[word], urls)
		}

		for i := range urls {
			wantUrl := urls[i]
			gotUrl := got[word][i]

			if strings.Compare(gotUrl, wantUrl) != 0 {
				t.Errorf("Urls don't match got = %v, want %v", gotUrl, wantUrl)
			}
		}

	}
}

func Test_containString(t *testing.T) {
	type args struct {
		words []string
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contain string",
			args: args{
				words: []string{
					"word1",
					"word2",
				},
				word: "word2",
			},
			want: true,
		},
		{
			name: "Does not contain string",
			args: args{
				words: []string{
					"word1",
					"word2",
				},
				word: "word3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containString(tt.args.words, tt.args.word); got != tt.want {
				t.Errorf("containString() = %v, want %v", got, tt.want)
			}
		})
	}
}
