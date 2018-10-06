package main

import (
	"fmt"
	"strings"
)

type myString struct {
	str string
}

type myString2 struct {
	myString
}

func (ms myString) printUpper() {
	fmt.Println(strings.ToUpper(ms.str))
}

// Will ovveride
func (ms myString2) printUpper() {
	fmt.Printf("PrintUpper from MyString2 type: %v\n", strings.ToUpper(ms.str))
}

func main() {
	ms2 := myString2{}

	ms2.str = "Ciao"

	ms2.printUpper()
}
