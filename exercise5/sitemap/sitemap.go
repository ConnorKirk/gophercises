package sitemap

import (
	"encoding/xml"
	"fmt"
	"gophercise/exercise4/linksearch"
	"io"
	"os"
)

var seen map[string]bool
var links []linksearch.Link

type urlSet struct {
	XMLName      xml.Name `xml:"urlset"`
	XMLNamespace string   `xml:"xmlns,attr`
	URLs         []string `xml:"url>loc"`
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
		u.URLs = append(u.URLs, k)
	}

	return u
}

func buildXMLOutput(u urlSet) {

	io.WriteString(os.Stdout, xml.Header)
	enc := xml.NewEncoder(os.Stdout)

	enc.Indent("  ", "    ")
	enc.Encode(u)

	enc.Flush()

}
