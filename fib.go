package main

import (
	"fmt"
)

func main() {
	fibonacciNumbers()
}

// fibonacciNumbers prints fibonacci numbers
// up to what an int holds
func fibonacciNumbers() {
	x := 1
	y := 1
	fmt.Printf("1\n1\n")

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
