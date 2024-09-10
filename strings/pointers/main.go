package main

import (
	"fmt"
	"unicode"
)

func isLatin(input string) bool {
	for _, letter := range input {
		if !unicode.Is(unicode.Latin, letter) {
			return false
		}
	}
	return true
}

func main() {
	name := "oleg"

	fmt.Println(isLatin(name))
}