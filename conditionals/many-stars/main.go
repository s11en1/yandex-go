package main

import (
	"fmt"
)

func main(){
	var starCount int

	fmt.Scanln(&starCount)

	if starCount < 0 {
		fmt.Println("Некорректный ввод")
	} else if starCount < 10 {
		fmt.Println("Число меньше 10")
	} else if starCount < 100 {
		fmt.Println("Число меньше 100")
	} else if starCount < 1000 {
		fmt.Println("Число меньше 1000")
	} else {
		fmt.Println("Число больше или равно 1000")
	}
}