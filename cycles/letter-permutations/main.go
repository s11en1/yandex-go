package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	factorial := 1

	for i := 2; i <= input; i++ {
		factorial *= i;
	}  

	fmt.Println(factorial)
}