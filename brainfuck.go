/*

i	instructions
ip	instruction pointer

b	byte array of length n = 4096; adjust as needed
bp	byte pointer

match	a counter of open (closed) brackets, used to find the matching open (closed) bracket

example programs:

Hello World!
"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

I LOVE YOU
"++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++."

AAA
"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.>-]"

ABC
"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.+>-]"

A
"+>++++++++[<++++++++>-]<."
The 2nd register acts as a counter, allowing us to run the loop 8 times. Inside the loop, we increment the 1st register by 8. When the counter reaches 0, the loop ends and we move the back to the print the value in the first register. ASCII 'A' is 1 + 8 * 8 = 65. Notice that 8 + 8 = 16. If we want to express 64 as a different product, we'd have to double the number of pluses in one place for every halving in the other: 16 + 4 = 20, 32 + 2 = 34. In general, the closer we can get both factors to the square root of the target quantity the better. (For rectangles of a given area, the square is the rectangle with the smallest perimeter.)

AAA
"+>++++++++[<++++++++>-]+++[<.>-]"

ABC
">++++++++[<++++++++>-]+++[<+.>-]"

123456789
">+++++++[<+++++++>-]+++++++++[<.+>-]"

0123456789
">+++++++<->[<+++++++>-]++++++++++[<.+>-]"

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
