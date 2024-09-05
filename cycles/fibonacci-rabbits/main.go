package main 

import (
	"fmt"
)

func main() {
	var input int

	fmt.Scanln(&input)

	temp, current, next, step := 0, 0, 1, 0
	
	for {
		if(step == 10) {
			break
		}

		if(current >= input) {
			fmt.Println(current)
			step++
		}

		temp = next
		next += current
		current = temp
	}
}