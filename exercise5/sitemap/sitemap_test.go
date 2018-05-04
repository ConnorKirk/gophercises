package sitemap

import (
	"testing"
)

func TestConvertLink(t *testing.T) {
	site := "www.example.com"
	relativeLink := []string{
		"/this-is-a-relative-link",
		"./about",
		"/does-this-work/",
		"./",
	}

	expected := []string{
		"www.example.com/this-is-a-relative-link",
		"www.example.com/about",
		"www.example.com/does-this-work",
		"www.example.com",
	}

	for i, l := range relativeLink {
		cl := convertLink(site, l)
		if cl != expected[i] {
			t.Errorf("Expected %v\nGot %v", expected[i], cl)
		}
	}

}

func TestBuildURLSet(t *testing.T) {
	seen = map[string]bool{
		"www.example.com":       true,
		"www.example.com/about": true,
		"www.example.com/faq":   true,
	}

	u := buildURLSet(seen)

	if len(u.URLs) != 3 {
		t.Errorf("Expecting 3, got %v", len(u.URLs))
	}
}

func TestBuildXMLOutput(t *testing.T) {
	u := urlSet{
		XMLNamespace : `xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"`,
		URLs :[]string{
			"www.example.com",
			"www.example.com/about",
		},
	}

	buildXMLOutput(u)
}
