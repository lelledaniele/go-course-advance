package main

import (
	"fmt"
)

func main() {
	hi := struct {
		English string
		Italian string
	}{"hi", "ciao"}

	fmt.Printf("%#v\n", hi)
}
