package main

import (
	"fmt"
	"regexp"
)

func main() {
	// r, err := regexp.Compile("[0-9]+")
	r := regexp.MustCompile("[0-9]+") // panic if regex is invalid

	// if err != nil {
	// 	fmt.Println("Error")
	// 	os.Exit(1)
	// }

	fmt.Printf("Does it match? %v\n", r.MatchString("Ciao 123"))
}
