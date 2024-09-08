package main

import (
	"fmt"
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 || nums == nil {
		return nil, errors.New("Something went wrong")
	}
	
	var result []int

	for _, value := range nums {
		if value >= limit {
			continue
		}

		if len(result) == n {
			break
		}

		result = append(result, value)
	}

	return result, nil
}

func main() {
	nums := []int{5,4,3,2,1}

	fmt.Println(UnderLimit(nums, 3, -1))
}