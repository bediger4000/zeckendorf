package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

// Action: what to occasionally do to stream of '1' and '0'
type Action int

const (
	Nothing Action = 0
	Flip    Action = iota
	Insert  Action = iota
	Drop    Action = iota
)

var action = []Action{
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Nothing, Nothing, Nothing, Nothing, Nothing, Nothing, Nothing,
	Flip, Insert, Drop,
}

func main() {
	flag.Parse()
	fin := os.Stdin
	if flag.NArg() > 1 {
		var err error
		if fin, err = os.Open(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		defer fin.Close()
	}

	max := len(action)

	buf := make([]byte, 1)

	var n int
	var err error

	for n, err = fin.Read(buf); err == nil && n == 1; n, err = fin.Read(buf) {
		switch action[rand.Intn(max)] {
		case Flip:
			switch buf[0] {
			case '1':
				buf[0] = '0'
			case '0':
				buf[0] = '1'
			}
		case Insert:
			var inserted [1]byte
			switch rand.Intn(2) {
			case 0:
				inserted[0] = '0'
			case 1:
				inserted[0] = '1'
			}
			writeByte(inserted[:])
		case Drop:
			continue
		}
		writeByte(buf)
	}

	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "loop ending error: %v\n", err)
		fmt.Fprintf(os.Stderr, "loop ending error, read %d bytes\n", n)
	}
}

func writeByte(buf []byte) {
	m, errW := os.Stdout.Write(buf)
	if m != 1 {
		fmt.Fprintf(os.Stderr, "wrote %d byte(s), should have written 1\n", m)
		return
	}
	if errW != nil {
		fmt.Fprintf(os.Stderr, "writing 1 byte error: %v\n", errW)
	}
}
