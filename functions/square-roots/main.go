package main

import (
	"fmt"
	"math"
)

func main() {
	SqRoots()
}

func SqRoots() {
	var a, b , c float64

	fmt.Scanln(&a, &b, &c)

	discriminant := findDiscriminant(a, b ,c)

	if discriminant < 0 {
		fmt.Println(0, 0)
	} else if discriminant == 0 {
		fmt.Println(-b / (2 * a))
	} else {
		firstRoot := (-b - math.Sqrt(discriminant)) / (2 * a)
		secondRoot := (-b + math.Sqrt(discriminant)) / (2 * a)

		fmt.Println(firstRoot, secondRoot)
	}
}

func findDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4.0 * a * c
}