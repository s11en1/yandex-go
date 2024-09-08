package main

import (
	"fmt"
)

func Join(nums1, nums2 []int) []int {
	newLength := len(nums1) + len(nums2)
	result := make([]int, 0, newLength)

	for _, value := range nums1 {
		result = append(result, value)
	}

	for _, value := range nums2 {
		result = append(result, value)
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}

	newSlice := Join(nums1, nums2)

	fmt.Println(newSlice)
}