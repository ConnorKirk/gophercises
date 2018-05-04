package sitemap

import (
	"fmt"
	"gophercise/exercise4/linksearch"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func traverse(page string, basesite string) {
	fmt.Printf("Traversing: %v\n", page)
	// Add to seen
	seen[page] = true

	s := getPage(page)
	links := linksearch.Search(strings.NewReader(s))
	for _, l := range links {
		nextLink := convertLink(basesite, l.Href)

		if seen[nextLink] || !sameSite(nextLink, page) {
			continue
		}
		traverse(nextLink, basesite)
	}

	//
}

func getPage(s string) string {
	resp, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return string(bodyBytes)
}

//convertLink converts a relative url to an absolute url
func convertLink(baseSite string, link string) string {
	//Check if relative or absolute link

	if strings.HasSuffix(link, "/") {
		link = link[:len(link)-1]
	}

	if strings.HasPrefix(link, ".") {
		return baseSite + link[1:]
	}

	if strings.HasPrefix(link, "/") {
		return baseSite + link
	}

	return link
}

func sameSite(link, site string) bool {
	return strings.HasPrefix(link, site)
}
