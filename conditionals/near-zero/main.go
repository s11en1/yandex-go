package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	if a == 0 || b == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else if a > 0 && b > 0 {
		fmt.Println("Оба числа положительные")
	} else if a < 0 && b < 0 {
		fmt.Println("Оба числа отрицательные")
	} else {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}