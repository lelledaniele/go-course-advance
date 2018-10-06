package main

import "fmt"

func returnVariableDeclaration(i []int) (n int) {
	return n
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(returnVariableDeclaration(ints))

Abort:
	for i, n := range ints {
		fmt.Println("Loop", i)

		switch {
		case n%2 == 0:
			continue
		case n == 5:
			break Abort // Useful for multiple for
			// continue FirstLoop
		default:
			fmt.Println(n)
		}
	}

	// Variable scope
	if true {
		ints := []int{10, 40}

		fmt.Println(ints)
	} else {
		// ints := []int{10, 40} // Error
	}

	// ints := []int{10, 40} // Error

	//

	n, err := fmt.Println("Ciao")
	// n, err := fmt.Println("Ciao 2") // Error

	fmt.Println(n, err)

	//

	n1, err := fmt.Println("Ciao")
	n2, err := fmt.Println("Ciao 2") // Fine

	fmt.Println(n1, n2, err)

	//

	// _, err := fmt.Println("Ciao")   // Error - No new variables
	// _, err := fmt.Println("Ciao 2") // Error - No new variables
	_, err = fmt.Println("Ciao 2") // Fine

	fmt.Println(err)
}
