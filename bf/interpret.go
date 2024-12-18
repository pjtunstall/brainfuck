package bf

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Interpret(i string, stopChan chan(struct{}), t time.Duration) (string, error) {
	n := 30_000
	b := make([]byte, n)	
	ip := 0 // instruction pointer
	bp := 0 // byte pointer

	var o strings.Builder
	var err error

	for ip < len(i) {
		select {
			case <-stopChan:
				return "", fmt.Errorf("timed out after %d seconds", t / time.Second)
			default:
			switch i[ip] {
			case '+':
				b[bp]++
			case ',':
				var input byte
				reader := bufio.NewReader(os.Stdin)
				input, err = reader.ReadByte()
				if err != nil {
					return "", err
				}
				b[bp] = input
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
	}
	
	return o.String(), nil
}
