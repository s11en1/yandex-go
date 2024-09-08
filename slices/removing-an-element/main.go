package main

import (
	"fmt"
)

func Clean(nums []int, x int) []int {
	move, sliceLength := 0, len(nums)

	for i := 0; i < sliceLength; i++ {
		if nums[i] == x {
			move++
		} else {
			nums[i - move] = nums[i]
		}
	}

	return nums[:sliceLength - move]
}

func main() {
	nums := []int{1,3,4,5,3,5,3,5,5}

	fmt.Println("Слайс до изменения:", nums)

	fmt.Println("Слайс после изменения:", Clean(nums,3))
}