package main

import (
	"fmt"
	"math/rand"
	"time"

	"theMatrix/pkg/console"
	"theMatrix/pkg/symbols"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// symbols variable contains all symbols to be displayed
	symbols := symbols.Symb()

	oneLine := make(chan string)
	consoleSize := make(chan int)

	for {

		go console.ConsoleSize(consoleSize)

		go lineGenerator(symbols, consoleSize, oneLine)
		time.Sleep(time.Millisecond * 700)
		fmt.Println(<-oneLine)
	}
}

// lineGenerator function creates a string of random symbols and spaces
func lineGenerator(symbols map[int]string, consoleSize chan int, oneLine chan string) {
	res := ""

	termSize := <-consoleSize

	// spaceNumFlag variable is the ratio of white spaces to symbols
	// here it is 6/10
	spaceNumFlag := (len(symbols) / 10) * 6

	// this loop creates a string of random symbols and white spaces between symbols in random places
	// with the size of terminal window
	for i := 0; i < termSize; i++ {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		rand := r1.Intn(len(symbols))

		if rand >= spaceNumFlag {
			res += " "
		} else {
			res += symbols[rand]
		}

	}

	oneLine <- res
}
