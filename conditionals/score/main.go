package main

import (
	"fmt"
)

func main() {
	var score int

	fmt.Println("Введите вашу оценку: ")
	fmt.Scanln(&score)

	if score >= 60 {
		fmt.Println("Отличная работа!")
	} else {
		fmt.Println("Нужно усерднее учиться!")
	}
}