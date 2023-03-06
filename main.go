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

	oneLine := make(chan string)
	consoleSize := make(chan int)

	for {

		go console.ConsoleSize(consoleSize)

		go lineGenerator(symbols, consoleSize, oneLine)
		time.Sleep(time.Second)
		fmt.Println(<-oneLine)
	}
}

// lineGenerator function creates a string of random symbols and spaces
func lineGenerator(symbols map[int]string, consoleSize chan int, oneLine chan string) {
	res := ""

	termSize := <-consoleSize

	// spaceNum variable contains the possible number of white spaces in a row of symbols
	spaceNum := termSize / 2

	// this loop creates a string of random symbols  with the size of terminal window
	for i := 0; i < termSize; i++ {

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
