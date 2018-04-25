package linksearch

import (
	"fmt"
)

//Link is a struct containing a hyperref and the displayed text
type Link struct {
	Href string
	Text string
}

func (l *Link) String() {
	fmt.Printf("Href: %v, Text: %v", l.Href, l.Text)
}
