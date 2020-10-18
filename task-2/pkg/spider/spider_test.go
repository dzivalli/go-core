package spider

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_pageText(t *testing.T) {
	doc := `<html><head></head><body><div>word1</div><span>word2 word3 word4</span><script>script</script></body></html>`
	reader := strings.NewReader(doc)
	page, _ := html.Parse(reader)

	got := pageText(page, []string{})
	want := []string{"word1", "word2 word3 word4"}

	if len(got) != len(want) {
		t.Errorf("Length does not match, got: %v, want: %v", got, want)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Wrong entry, got: %v, want: %v", got[i], want[i])
		}
	}

}
