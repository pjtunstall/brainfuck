/*

`i`		instructions
`ip`	instruction pointer

`b`		byte array of length n = 4096; adjust as needed
`bp`	byte pointer

`match`	a counter of open (closed) brackets, used to find the matching open (closed) bracket

*/

package main

import (
	"bufio"
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

	i := os.Args[1]
	ip := 0

	n := 4096
	b := make([]byte, n)

	bp := 0

	for ip < len(i) {
		switch i[ip] {
		case '+':
			b[bp]++
		case ',':
			var input byte
			fmt.Scanf("%c", &input)
			b[bp] = input
			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
		case '-':
			b[bp]--
		case '.':
			fmt.Print(string(b[bp]))
		case '<':
			bp--
			if bp < 0 {
				bp = n - 1
			}
		case '>':
			bp++
			if bp == n {
				bp = 0
			}
		case '[':
			if b[bp] == 0 {
				ip = openBracket(i, ip)
			}
		case ']':
			if b[bp] != 0 {
				ip = closeBracket(i, ip)
			}
		}
		ip++
	}
	fmt.Println()
}

func openBracket(i string, ip int) int {
	match := 1
	for match > 0 {
		ip++
		if ip >= len(i) {
			fmt.Println("\nerror: instruction pointer out of range to the right.")
			os.Exit(1)
		}
		if i[ip] == 91 {
			match++
		}
		if i[ip] == 93 {
			match--
		}
	}
	return ip
}

func closeBracket(i string, ip int) int {
	match := 1
	for match > 0 {
		ip--
		if ip < 0 {
			fmt.Println("\nerror: instruction pointer out of range to the left.")
			os.Exit(1)
		}
		if i[ip] == 93 {
			match++
		}
		if i[ip] == 91 {
			match--
		}
	}
	return ip
}
