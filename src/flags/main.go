package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.Int("version", 1, "Get version")
	flag.Parse()

	fmt.Println(*version)
}

// go run main.go -version=2
