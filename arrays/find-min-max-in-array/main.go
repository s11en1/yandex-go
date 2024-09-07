package main

import (
	"fmt"
)

func FindMinMaxInArray(array [10]int) (int, int) {
	minIndex, maxIndex := 0, 0

	for i, _ := range array {
		if array[i] > array[maxIndex] {
			maxIndex = i
		}

		if array[i] < array[minIndex] {
			minIndex = i
		}
	}

	return array[maxIndex], array[minIndex]
}

func main() {
	min, max := FindMinMaxInArray([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Println(min, max)
}