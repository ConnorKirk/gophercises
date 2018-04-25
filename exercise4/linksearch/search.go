package linksearch

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

//Links is an array of link structs
var links []Link

// Search takes an ioReader and performs a dfs of a the html doc. It returns an array of found links
func Search(page io.Reader) []Link {
	doc, err := html.Parse(page)

	if err != nil {
		log.Fatalf("couldn't parse %v: %v", page, err)
	}

	traverse(doc)

	return links
}

// traverse performs a dfs looking for <a> tags. It appends found links to the global var Links
func traverse(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		foundLink := Link{}
		for _, a := range n.Attr {
			if a.Key == "href" {
				foundLink.Href = a.Val
				foundLink.Text = getData(n)
			}

			links = append(links, foundLink)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}

// getData initiates the recursive function textTraverse and returns the final string of the <a> tag provided
func getData(n *html.Node) string {

	var ret string

	ret = textTraverse(n, "")

	return ret
}

// textTraverse concatenates all of the text of a provided <a> node
func textTraverse(n *html.Node, s string) string {
	var ret string
	if n.Type == html.TextNode {
		ret += n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		ret += textTraverse(c, s)

	}

	return ret

}
