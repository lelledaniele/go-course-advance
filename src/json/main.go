package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	inputJSON := []byte(`[
		{"eng": "hi", "ita": "ciao"},
		{"eng": "hi", "ita": "salve"}
	]`)
	var class []struct {
		eng string
		ita string
	}

	_ = json.Unmarshal(inputJSON, &class)

	fmt.Println(class)
}
