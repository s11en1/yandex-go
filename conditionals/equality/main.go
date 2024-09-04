package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	numsCount, _ := fmt.Scanln(&a, &b, &c)

	if numsCount < 3 {
		fmt.Println("Некорректный ввод")
	} else if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if a == b || b == c || c == a {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}