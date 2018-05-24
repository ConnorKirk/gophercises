package main

import "fmt"

func main() {
	var length int
	var input string
	var k int

	fmt.Scanf("%v\n%s\n%v", &length, &input, &k)

	// For each rune in input string, shift by k % 24
	fmt.Println(length, input, k)
	var out []rune
	for _, r := range input {
		new := (int(r) + k)
		out = append(out, rune(new))
	}
	fmt.Printf("%q %v", out, out)
}
