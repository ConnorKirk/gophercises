package sitemap

import (
	"fmt"
	"gophercise/exercise4/linksearch"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var seen map[string]bool
var links []linksearch.Link

// Build takes a string declaring the site to be built
func Build(site string) (string, error) {

	fmt.Printf("Building a site map for %v\n", site)
	seen = make(map[string]bool)

	//initiate traverse
	traverse(site)

	// convert seen to xml

	// return xml string

	return "", nil
}

func traverse(page string) {
	fmt.Printf("Traversing: %v\n", page)
	// Add to seen
	seen[page] = true
	resp, err := http.Get(page)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("%v", string(bodyBytes))
	links := linksearch.Search(resp.Body)
	fmt.Printf("Found %v links", len(links))
	for _, l := range links {
		wholeLink := convertLink(page, l.Href)

		if !seen[wholeLink] {
			continue
		}
		traverse(wholeLink)
	}

	//
}

//convertLink converts a relative url to an absolute url
func convertLink(site string, link string) string {
	//Check if relative or absolute link
	if strings.HasPrefix(link, "/") {
		return site + link
	}

	return link
}
