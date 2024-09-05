package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	for i := 3; i <= input; i += 3 {
		fmt.Println(i)
	}
}