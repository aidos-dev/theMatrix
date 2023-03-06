package main

import (
	"fmt"

	"theMatrix/console"
	"theMatrix/symbols"
)

func main() {
	symbols := symbols.Symb()
	fmt.Println(symbols)

	consoleSize := console.ConsoleSize()
	fmt.Printf("console size: %v\n", consoleSize)
}
