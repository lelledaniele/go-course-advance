package main

import (
	"fmt"
)

// Buffer example
type Buffer struct{}

func (Buffer) process() {
	fmt.Println("Process")
}

func (Buffer) get() {
	fmt.Println("Get")
}

var available = make(chan Buffer, 10)
var toController = make(chan Buffer)

func worker() {
	for {
		var b Buffer

		select {
		case b = <-available:
		default:
			b = Buffer{}
		}

		b.get()
		toController <- b
	}
}

func controller() {
	for {
		b := <-toController
		b.process()

		select {
		case available <- b:
		default:
		}
	}
}

func main() {
	var a []int
	a = make([]int, 20, 22)

	var c *[]int
	c = new([]int)

	fmt.Printf("%d\n", len(a))
	fmt.Printf("%d\n", len(*c))

	go controller()
	go worker()

	// ???
}
