package main

import (
	"fmt"
)

func StringLength(input string) int {
	return len([]rune(input))
}

func main() {
	fmt.Println(StringLength("Эйяфьядлайёкюдль"))
}