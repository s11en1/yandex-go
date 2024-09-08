package main

func SliceCopy(nums []int) []int {
	result := make([]int, len(nums), len(nums))

	for index, value := range nums {
		result[index] = value
	}

	return result
}