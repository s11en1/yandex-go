package main

import (
	"fmt"
)

func FiveSteps(array [5]int) [5]int {
	var newArray [5]int

	for i, j := range array {
		newArray[4 - i] = j
	}

	return newArray
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(FiveSteps(array))
}