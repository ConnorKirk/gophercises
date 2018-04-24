package linksearch

import (
	"fmt"
)

type Link struct {
	Href string
	Text string
}

func (l *Link) String() {
	fmt.Printf("Href: %v, Text: %v", l.Href, l.Text)
}
