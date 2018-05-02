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

	// absoluteLink := "/this-is-an-absolute-link"
	// l := convertLink(site, relativeLink)
	// if l != "www.example.com/this-is-a-relative-link" {
	// 	t.Errorf("relative link not converted. Got %v", l)
	// }

	// al := convertLink(site, absoluteLink)
	// if al != "www.example.com/this-is-an-absolute-link" {
	// 	t.Errorf("relative link not converted. Got %v", al)
	// }

}
