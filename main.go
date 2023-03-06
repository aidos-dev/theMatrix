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

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// spaceNum variable contains the possible number of white spaces in a row of symbols
	spaceNum := termSize / 10

	// this loop creates a string of random symbols  with the size of terminal window
	for i := 0; i < termSize; i++ {

		rand := r1.Intn(len(symbols))

		res += symbols[rand]

	}

	// this loop replaces some symbols with white spaces in random places
	for i := 0; i < spaceNum; i++ {

		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)

		spaceRand := r2.Intn(termSize)

		res = strings.Replace(res, string(res[spaceRand]), " ", 1)
	}

	oneLine <- res
}
