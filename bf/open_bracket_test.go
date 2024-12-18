package bf

import (
	"errors"
	"testing"
)

func TestOpenBracket_OutOfRangeError(t *testing.T) {
	input := "["

	_, err := openBracket(input, 0)
	if err == nil {
		t.Errorf("Expected an error, but got: nil")
	}

	var outOfRangeErr OutOfRangeError
	if !errors.As(err, &outOfRangeErr) {
		t.Errorf("Expected OutOfRangeError, but got: %T", err)
	}

	if outOfRangeErr.IP != 1 {
		t.Errorf("Expected instruction pointer to be 1, but got %d", outOfRangeErr.IP)
	}
}

func TestOpenBracket_ValidInput(t *testing.T) {
	input := "[[]]"
	ip, err := openBracket(input, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if ip != 3 {
		t.Errorf("Expected instruction pointer to be 3, but got %d", ip)
	}
}
