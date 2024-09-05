package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	sum := 0

	for i := 1; i <= input; i++ {
		sum += i
	}

	fmt.Println(sum)
}