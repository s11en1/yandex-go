package main

import (
	"fmt"
)

func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}

func main() {
	firstName := "Олег"
	secondName := "Галкин"

	fmt.Println(ConcatenateStrings(firstName, secondName))
}
