package main

import (
	"fmt"
)

func SumOfArray(array [6]int) int {
	sum := 0

	for _, value := range array {
		sum += value
	}

	return sum
}

func main() {
	var array = [...]int{1, 2, 3, 4, 5, 6}

	sum := SumOfArray(array)

	fmt.Println(sum)
}