package bf

import (
	"math/rand"
	"testing"
)

func TestWrappedInterpret_HappyPath(t *testing.T) {
	s, err := WrappedInterpret("++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++.", 1)
	if err == nil {
		t.Errorf("expected a timeout error, but got: nil")
	}
	if s != "" {
		t.Errorf("expected \"\", but got %q", s)
	}
}

// A fuzz test to check for unexpected panics.
func TestWrappedInterpret_NoPanic(t *testing.T) {
	defer func() {
        if r := recover(); r != nil {
            t.Fatalf("Test panicked unexpectedly: %v", r)
        }
    }()

	bfStrings := generateRandomBrainfuck(256, 16)
	for _, bfString := range bfStrings {
		_, _ = WrappedInterpret(bfString, 3)
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