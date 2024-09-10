package main

import (
	"fmt"
)

func FindMaxKey(m map[int]int) int {
	var max int
	firstKey := true

	for key := range m {
		if firstKey {
			max = key
			firstKey = false
			continue
		}

		if key > max {
			max = key
		}
	}
	return max
}

func main() {
	m := map[int]int{
		-7: 1,
		-1: 38,
	}
	fmt.Println(FindMaxKey(m))
}
