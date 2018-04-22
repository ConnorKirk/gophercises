package main

import (
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

var data = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.

func main() {
	pathMap, err := parseYAML([]byte(data))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(pathMap)
}

type ParsedYaml struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(yml []byte) (map[string]string, error) {
	var paths []ParsedYaml

	err := yaml.Unmarshal([]byte(yml), &paths)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	pathMap := make(map[string]string)

	for _, path := range paths {
		pathMap[path.Path] = path.Url
	}
	return pathMap, nil
}
