package main

import (
	"brainfuck/bf"
	"fmt"
	"os"
)

func main() {
	input, timeout, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		usage()
		os.Exit(1)
	}

	output, err := bf.WrappedInterpretWithTimeout(input, timeout)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", output)
	}
}

func usage() {
	fmt.Println("Usage: go run . \"Brainfuck instructions\" [TIMEOUT] (in seconds)")
		fmt.Println("\nExamples:")
		fmt.Println("  go run . \">+++++++<->[<+++++++>-]++++++++++[<.+>-]\"")
		fmt.Println("  go run . \"+[]\" 1 # infinite loop, timeout after 1 second")
		fmt.Println()
}
