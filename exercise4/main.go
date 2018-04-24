package main

import (
	"flag"
	"fmt"
	"gophercise/exercise4/linksearch"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	page := flag.String("page", "ex1.html", "page you wish to search")
	flag.Parse()

	file, err := ioutil.ReadFile(*page)
	if err != nil {
		log.Fatal(err)
	}

	read := strings.NewReader(string(file))
	links := linksearch.Search(read)

	fmt.Println("Found links:", links)

}
