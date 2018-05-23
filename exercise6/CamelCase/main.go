package main

import (
	"fmt"
	"strings"
)

func main() {
	// Take input from stdin

	var input string
	fmt.Scanf("%s\n", &input)
	//fmt.Printf("Input is %v\n", input)

	fmt.Println(countCamelCase(input))

	// Split word on Capitals
	// Count words

}

func countCamelCase(input string) int {
	answer := 1

	for _, r := range input {
		str := string(r)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	return answer
}
