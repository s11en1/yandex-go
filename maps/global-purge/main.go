package main

func DeleteLongKeys(m map[string]int) map[string]int {
	cleanedMap := map[string]int{}

	for key := range m {
		if len([]rune(key)) >= 6 {
			cleanedMap[key] = m[key]
		}
	}

	return cleanedMap
}
