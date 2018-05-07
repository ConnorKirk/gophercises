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

	expected := `<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
		<url>
			<loc>www.example.com</loc>
		</url>
		<url>
			<loc>www.example.com/about</loc>
		</url>
	</urlset>
	`
	u := urlSet{
		URLs: []URL{
			URL{("www.example.com")},
			URL{("www.example.com/about")},
		},
	}

	out := buildXMLOutput(u)

	if out != expected {
		t.Errorf("Output differs to expected. Expected:\n%v,\n Got:\n%v", expected, out)
	}
}
