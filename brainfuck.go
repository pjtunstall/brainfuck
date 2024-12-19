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

	output, err := bf.WrappedInterpret(input, timeout)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", output)
	}
}
