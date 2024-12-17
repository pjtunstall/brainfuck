package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*

`i`		instructions
`ip`	instruction pointer

`b`		byte array of length n = 4096; adjust as needed
`bp`	byte pointer

`match`	a counter of open (closed) brackets, used to find the matching open (closed) bracket

*/

type OutOfRangeError struct {
	IP int
	Msg string
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("Error: %s (instruction pointer: %d)", e.Msg, e.IP)
}

func interpret(i string) (string, *OutOfRangeError) {
	ip := 0

	n := 4096
	b := make([]byte, n)

	bp := 0

	var o strings.Builder
	var err *OutOfRangeError

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
			o.WriteByte(b[bp])
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
				ip, err = openBracket(i, ip)
				if err != nil {
					return "", err
				}
			}
		case ']':
			if b[bp] != 0 {
				ip, err = closeBracket(i, ip)
				if err != nil {
					return "", err
				}
			}
		}
		ip++
	}
	
	return o.String(), nil
}
