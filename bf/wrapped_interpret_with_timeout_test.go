package bf

import (
	"math/rand"
	"strings"
	"testing"
)

func TestWrappedInterpretWithTimeout(t *testing.T) {
	s, err := WrappedInterpretWithTimeout("++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++.", 1)
	if err == nil {
		t.Errorf("Expected a timeout error, but got: nil")
	}
	if s != "" {
		t.Errorf("Expected '', but got '%q'", s)
	}

	// A fuzz test to check for unexpected panics.
	bfStrings := generateRandomBrainfuck(32, 16)
	for _, bfString := range bfStrings {
		_, err := WrappedInterpretWithTimeout(bfString, 3)
		if err != nil {
			if strings.Split(err.Error(), " ")[0] != "timed" {
				t.Errorf("Error: %s", err)
			}
		}
	}
}

func generateRandomBrainfuck(n int, length int) []string {
	symbols := []rune{'+', '-', '.', '<', '>', '[', ']'}
	var result []string

	for i := 0; i < n; i++ {
		var bfString []rune
		for j := 0; j < length; j++ {
			bfString = append(bfString, symbols[rand.Intn(len(symbols))])
		}
		result = append(result, string(bfString))
	}

	return result
}