package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := map[string]string{}

	for key := range m {
		newMap[m[key]] = key
	}

	return newMap
}

func main() {
	m := map[string]string{
		"Олег": "19 лет",
		"Эд":   "20 лет",
	}

	fmt.Println(SwapKeysAndValues(m))
}
