package main

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N == 1 {
		fmt.Println("YES")
		return
	} 
	
	if N % 2 == 0 {
		IsPowerOfTwoRecursive(N / 2)
	} else {
		fmt.Println("NO")
		return
	}
}

func main() {
	var n int

	fmt.Scanln(&n)

	IsPowerOfTwoRecursive(n)
}