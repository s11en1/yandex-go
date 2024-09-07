package main

import (
	"fmt"
)

func SumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}

	return n % 10 + SumDigitsRecursive(n / 10)
}

func main() {
	sum := SumDigitsRecursive(123)

	fmt.Println(sum)
}