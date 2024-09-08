package main

import (
	"fmt"
)

func Mix(nums []int) []int {
	length := len(nums)
	n := len(nums) / 2
	index := 0
	
	mixedSlice := make([]int, length, length)

	for i := 0; i < n ; i++ {
		mixedSlice[index], mixedSlice[index + 1] = nums[i], nums[i + n]
		index += 2
	}

	return mixedSlice
}

func main() {
	nums := []int{1,3,5,2,4,6}

	mixed := Mix(nums)

	fmt.Println(mixed)
}

// index = 0 | index + 1 = 1 | i = 0 | i + n = 3 
// index = 2 | index + 1 = 3 | i = 1 | i + n = 4 
// index = 4 | index + 1 = 5 | i = 2 | i + n = 5