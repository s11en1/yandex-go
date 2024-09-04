package main

import (
	"fmt"
)

func main() {
	var pageCount int

	var response string = "Число "

	fmt.Scanln(&pageCount)

	if pageCount > 0 {
		response += "положительное и "
	} else if pageCount < 0 {
		response += "отрицательное и "
	} else {
		fmt.Println("Число равно нулю")
		return
	}

	if pageCount % 2 == 0 {
		response += "четное"
	} else {
		response += "нечетное"
	}

	fmt.Println(response)
}