package cyoa

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

func StoryHandler(s Story, startArc *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.Replace(path, "/", "", 1)
		if path == "" {
			path = *startArc
		}
		tmpl := template.Must(template.ParseFiles("cyoa/templates.html"))
		tmpl.Execute(w, s[path])
	}
}

func IndexHandler(s Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get list of keys from map
		var chapters []string
		for k := range s {
			err := append(chapters, k)
			if err != nil {
				fmt.Fprint(w, err)
			}
		}
		// Arrange alphabetically
		sort.Strings(chapters)
		// Print to screen
		fmt.Fprint(w, chapters)
	}
}
