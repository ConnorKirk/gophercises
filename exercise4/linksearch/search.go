package linksearch

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

func Search(page io.Reader) []Link {
	f, err := html.Parse(page)

	if err != nil {
		log.Fatalf("couldn't parse %v: %v", page, err)
	}
	fmt.Println(f)

	return nil
}

func traverse(n *html.Node) []Link {
	fmt.Print(n.Data)

	return nil
}
