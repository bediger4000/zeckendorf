package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// fibs has Fibonacci less than or equal to n
	fibs := fibonacciNumbers(n)

	// The non-conseuctive fibonacci numbers that sum to n
	var digits []uint64

	m := len(fibs) - 1
	leftover := n - fibs[m]
	digits = append(digits, fibs[m])
	m -= 2 // nonconsecutive

	for leftover > 0 {
		if fibs[m] > leftover {
			m--
			continue
		}
		leftover = leftover - fibs[m]
		digits = append(digits, fibs[m])
		m -= 2 // nonconsecutive
	}

	// Check the digits, print them, and sum them
	var sum uint64
	spacer := ""
	for i := range digits {
		fmt.Printf("%s%d", spacer, digits[i])
		sum += digits[i]
		spacer = ", "
	}
	fmt.Printf("\nsum %d\n", sum)
	if sum != n {
		fmt.Print("Something's wrong")
	}
}

// fibonacciNumbers returns a slice of uint64 numbers,
// the fibonacci numbers less than argument n
func fibonacciNumbers(n uint64) []uint64 {
	ary := make([]uint64, 2)
	ary[0] = 1
	ary[1] = 1

	for i, j := 0, 1; ary[j] < n; i, j = j, j+1 {
		nxt := ary[i] + ary[j]
		if nxt > n {
			break
		}
		ary = append(ary, nxt)
	}

	return ary
}
