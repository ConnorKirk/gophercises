package main

import (
	"flag"
	"fmt"
	"gophercise/exercise3/cyoa"
	"log"
	"net/http"
	"os"
)

const port = ":8080"

func main() {
	filename := flag.String("file", "gopher.json", "Story to use")
	startArc := flag.String("start", "intro", "Arc to start")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	// Read and parse json file
	story, err := cyoa.GetStoryline(f)
	if err != nil {
		log.Fatalln(err)
	}
	m := http.NewServeMux()
	storyH := cyoa.StoryHandler(story, startArc)
	indexH := cyoa.IndexHandler(story)
	m.HandleFunc("/", storyH)
	m.HandleFunc("/index", indexH)
	m.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Listening on port ", port)
	http.ListenAndServe(port, m)

}
