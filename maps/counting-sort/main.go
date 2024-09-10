package main

func CountingSort(contacts []string) map[string]int {
	m := map[string]int{}

	for _, key := range contacts {
		m[key]++
	}

	return m
}
