package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	sum := 0

	for i := 1; i <= input; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			continue
		}

		sum += i
	}

	fmt.Println(sum)
}