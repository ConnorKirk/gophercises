package sitemap

import (
	"testing"
)

func TestConvertLink(t *testing.T) {
	site := "www.example.com"
	relativeLink := "/this-is-a-relative-link"
	absoluteLink := "/this-is-an-absolute-link"
	l := convertLink(site, relativeLink)
	if l != "www.example.com/this-is-a-relative-link" {
		t.Errorf("relative link not converted. Got %v", l)
	}

	al := convertLink(site, absoluteLink)
	if al != "www.example.com/this-is-an-absolute-link" {
		t.Errorf("relative link not converted. Got %v", al)
	}

}
