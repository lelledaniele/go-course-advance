package main

import (
	"fmt"
)

func incN(n int) int {
	return n + 1
}

func deferTest() {
	n := 1
	defer fmt.Printf("First - N is %d \n", n)

	defer fmt.Printf("Second - N is %d \n", incN(n))

	n = n + 1
	defer fmt.Printf("Third - N is %d \n", n)
}

func main() {
	deferTest()

	// Third - N is 2
	// Second - N is 2
	// First - N is 1
}
