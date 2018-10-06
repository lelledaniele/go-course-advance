package main

import "fmt"

func main() {
	c := make(chan bool)
	ints := []int{2, 3, 5}

	go func() {
		ints[1] = 50
		c <- true
	}()

	ints[1] = 60
	<-c
	for k, v := range ints {
		fmt.Println(k, v)
	}
}

// go run -race main.go
// ==================
// WARNING: DATA RACE
// Write at 0x00c4200a4008 by goroutine 6:
//   main.main.func1()
//       /home/lelle/projects/goadvance/src/race/main.go:10 +0x47

// Previous write at 0x00c4200a4008 by main goroutine:
//   main.main()
//       /home/lelle/projects/goadvance/src/race/main.go:14 +0xf3

// Goroutine 6 (running) created at:
//   main.main()
//       /home/lelle/projects/goadvance/src/race/main.go:9 +0xe1
// ==================
// 0 2
// 1 50
// 2 5
// Found 1 data race(s)
// exit status 66
