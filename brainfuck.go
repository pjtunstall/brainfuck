package main

import (
	"fmt"
	"os"
)

func main() {
	usage := "Usage: `go run . \"Insert Brainfuck instructions here!\"`"

	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments.")
		fmt.Println(usage)
		os.Exit(1)
	}

	output, err := interpret(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(output)
}
