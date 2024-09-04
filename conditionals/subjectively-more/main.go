package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	if a > b {
		fmt.Println("Первое число больше второго")
	} else if b > a {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}