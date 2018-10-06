package main

import (
	"errors"
	"fmt"
	"time"
)

type myError struct {
	str  string
	time time.Time
}

func (e myError) Error() string {
	return ""
}

func (e myError) print() {
	fmt.Printf("Err: %v at %d\n", e.str, e.time.Unix())
}

func main() {
	err := errors.New("Error 50")

	fmt.Printf("Err is %v\n", err)

	err2 := myError{
		str:  "Line 50",
		time: time.Now(),
	}

	// err := myError{
	// err.print() won't work
	err2.print()
}
