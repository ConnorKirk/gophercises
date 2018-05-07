package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"gophercise/exercise4/linksearch"
	"io"
)

var seen map[string]bool
var links []linksearch.Link

type URL struct {
	Address string `xml:"loc"`
}
type urlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	URLs    []URL    `xml:"url"`
}

// Build takes a string declaring the site to be built
func Build(site string) (string, error) {

	fmt.Printf("Building a site map for %v\n", site)
	seen = make(map[string]bool)

	//initiate traverse
	traverse(site, site)
	urlSet := buildURLSet(seen)
	fmt.Printf("Found %v links\n", len(seen))

	// convert seen to xml
	out, err := xml.Marshal(urlSet)
	if err != nil {
		return "", err
	}
	// return xml string
	return string(out), nil
}

func buildURLSet(seen map[string]bool) urlSet {
	var u urlSet

	for k := range seen {
		newURL := URL{k}
		u.URLs = append(u.URLs, newURL)
	}

	return u
}

func buildXMLOutput(u urlSet) string {
	b := new(bytes.Buffer)
	io.WriteString(b, xml.Header)
	enc := xml.NewEncoder(b)

	enc.Indent("  ", "    ")
	enc.Encode(u)

	enc.Flush()
	return b.String()
}
