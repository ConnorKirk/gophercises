package linksearch

import (
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	var p string

	p = `<html>
			<body>
				<h1>Hello!</h1>
				<a href="/other-page">A link to another page</a>
			</body>
		</html>`

	foundLinks := Search(strings.NewReader(p))

	if len(foundLinks) != 1 {
		t.Errorf("Expected 1 link, got %v", len(foundLinks))
	}

	if foundLinks[0].Href != "/other-page" {
		t.Errorf("Expected /other-page, got %v", foundLinks[0].Href)
	}
}
