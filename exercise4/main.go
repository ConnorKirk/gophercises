package main

import (
	"flag"
	"fmt"
	"gophercise/exercise4/linksearch"
	"net/http"
)

func main() {
	page := flag.String("page", "", "page you wish to search")
	res := http.Get(page)

	links := linksearch.Search(res.Body)

	fmt.Println(links)

}
