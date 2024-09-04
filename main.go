package main

import (
	"fmt"
	"golox/repl"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}

// func runFile(path string) {

// }
