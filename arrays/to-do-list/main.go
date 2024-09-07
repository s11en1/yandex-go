package main

import (
	"fmt"
)

func  PrettyArrayOutput(array [9]string) {
	boundaryIndex := len(array) - 3

	for index, value := range array {
		if index <= boundaryIndex {
			fmt.Println(index + 1, "я уже сделал:", value)
		} else {
			fmt.Println(index + 1, "не успел сделать:", value)
		}
	}
}