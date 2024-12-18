package main

import "fmt"

func usage() {
	fmt.Println("Usage: go run . \"Brainfuck instructions\" [TIMEOUT] (in seconds)")
		fmt.Println("\nExamples:")
		fmt.Println("  go run . \">+++++++<->[<+++++++>-]++++++++++[<.+>-]\"")
		fmt.Println("  go run . \"+[]\" 1 # infinite loop, timeout after 1 second")
		fmt.Println()
}