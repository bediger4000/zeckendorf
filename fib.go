package main

import (
	"fmt"
)

func main() {
	fibonacciNumbers()
}

// fibonacciNumbers prints fibonacci numbers
// up to what a uint64 holds
func fibonacciNumbers() {
	var x uint64
	var y uint64
	x, y = 1, 1
	fmt.Printf("1,\n1,\n")

	for {
		nxt := x + y
		if nxt < y {
			// overflow
			break
		}
		fmt.Printf("%d,\n", nxt)
		x, y = y, nxt
	}
}
