package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	if input < 0 {
		fmt.Println("Некорректный ввод")
		return
	}

	sum := 0

	for i := 1; i <= input; i += 2 {
		sum += i
	}

	fmt.Println(sum)
}