package main

import (
	"brainfuck/bf"
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Println("Usage: go run . \"Brainfuck instructions\" [TIMEOUT] (in seconds)")
		fmt.Println("\nExamples:")
		fmt.Println("  go run . \">+++++++<->[<+++++++>-]++++++++++[<.+>-]\"")
		fmt.Println("  go run . \"+[]\" 1")
		fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments")
		usage()
		os.Exit(1)
	}

	if len(os.Args) > 3 {
		fmt.Println("Error: too many arguments")
		usage()
		os.Exit(1)
	}

	timeout := 60 * 60 * time.Second
	var err error
	if len(os.Args) == 3 {
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			usage()
			os.Exit(1)
		}
		timeout = time.Duration(n * 1_000_000_000)
	}

	if timeout < 0 {
		fmt.Println("Error: timeout must be positive")
		usage()
		os.Exit(1)
	}

	output, err := bf.WrappedInterpretWithTimeout(os.Args[1], timeout)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", output)
	}
}
