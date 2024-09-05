package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	for i := 1; i <= 10; i++ {
		fmt.Println(input, "*", i, "=", input * i)
	}
}