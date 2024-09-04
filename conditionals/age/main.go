package main

import (
	"fmt"
)

func main() {
	var age int

	fmt.Println("Введите ваш возраст: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age < 65 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Труха")
	}
}