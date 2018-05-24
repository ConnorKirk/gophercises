package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var length int
	var input string
	var k int

	fmt.Scanf("%v\n%s\n%v", &length, &input, &k)

	// For each rune in input string, shift by k % 24

	var out []string
	for _, r := range input {
		//If a letter, apply cipher
		l := cipher(r, k)
		out = append(out, l)

	}

	fmt.Println(strings.Join(out, ""))

}

func cipher(r rune, shift int) string {
	if unicode.IsLower(r) {
		r = rune(((int(r) - 'a' + shift) % 26) + 'a')
	}

	if unicode.IsUpper(r) {
		r = rune(((int(r) - 'A' + shift) % 26) + 'A')
	}
	return string(r)
}
