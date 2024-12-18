package bf

import "fmt"

type OutOfRangeError struct {
	IP int
	Msg string
}

func (e OutOfRangeError) Error() string {
	return fmt.Sprintf("Error: %s (instruction pointer: %d)", e.Msg, e.IP)
}