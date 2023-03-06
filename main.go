package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"theMatrix/pkg/console"
	"theMatrix/pkg/symbols"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// symbols variable contains all symbols to be displayed
	symbols := symbols.Symb()
	// consoleSize function returns the current width of the terminal window
	consoleSize := console.ConsoleSize()
	fmt.Printf("console size: %v\n", consoleSize)

	oneLine := make(chan string)

	for {
		go lineGenerator(symbols, consoleSize, oneLine)
		time.Sleep(time.Second)
		fmt.Println(<-oneLine)
	}
}

func lineGenerator(symbols map[int]string, consoleSize int, oneLine chan string) {
	res := ""
	// spaceNum variable contains the possible number of white spaces in a row of symbols
	spaceNum := consoleSize / 2

	// this loop creates a string of random symbols  with the size of terminal window
	for i := 0; i < consoleSize; i++ {

		rand := rand.Intn(len(symbols))

		res += symbols[rand]

	}

	// this loop replaces some symbols with white spaces in random places
	for i := 0; i < spaceNum; i++ {
		spaceRand := rand.Intn(len(symbols))
		res = strings.Replace(res, string(res[spaceRand]), " ", 1)
	}

	oneLine <- res
}
