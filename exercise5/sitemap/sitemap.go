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

	for s := range seen {
		fmt.Printf("Seen: %v", s)
	}

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
	s := string(bodyBytes)
	links := linksearch.Search(strings.NewReader(s))
	fmt.Printf("Found %v links\n", len(links))
	for _, l := range links {
		wholeLink := convertLink(page, l.Href)

		if seen[wholeLink] || !sameSite(wholeLink, page) {
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

	if strings.HasPrefix(link, "./") {
		return site + link[1:]
	}

	return link
}

func sameSite(link, site string) bool {
	return strings.HasPrefix(link, site)
}
