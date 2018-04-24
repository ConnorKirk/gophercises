package linksearch

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

var Links []Link

func Search(page io.Reader) []Link {
	doc, err := html.Parse(page)

	if err != nil {
		log.Fatalf("couldn't parse %v: %v", page, err)
	}

	traverse(doc)

	return Links
}

func traverse(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		foundLink := Link{}
		for _, a := range n.Attr {
			if a.Key == "href" {
				foundLink.Href = a.Val
				foundLink.Text = n.FirstChild.Data
			}

			Links = append(Links, foundLink)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}
