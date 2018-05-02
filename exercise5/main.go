package main

import (
	"flag"
	"fmt"
	"gophercise/exercise5/sitemap"
	"log"
)

func main() {
	//Assumes that passed site is base site
	//Assumes protocol (http/s) is included in site
	// Read flag
	site := flag.String("site", "https://www.connorkirkpatrick.com", "site to build map of")
	flag.Parse()
	// Get html
	sm, err := sitemap.Build(*site)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(sm)
	// Start sitemap builder
}
