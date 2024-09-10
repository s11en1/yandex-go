package main

func SumOfValuesInMap(m map[int]int) int {
	sum := 0

	for key := range m {
		sum += m[key]
	}

	return sum
}
