package bf

import (
	"errors"
	"testing"
)

func TestCloseBracket_OutOfRangeError(t *testing.T) {
	input := "+]"

	_, err := closeBracket(input, 1)
	if err == nil {
		t.Errorf("expected an error, but got: nil")
	}	

	var outOfRangeErr OutOfRangeError	
	if !errors.As(err, &outOfRangeErr) {
		t.Errorf("expected OutOfRangeError, but got: %T", err)
	}

	if outOfRangeErr.IP != -1 {
		t.Errorf("expected instruction pointer to be -1, but got %d", outOfRangeErr.IP)
	}
}

func TestCloseBracket_ValidInput(t *testing.T) {
	input := "+>++++++++[<++++++++>-]<."
	ip, err := closeBracket(input, len(input) - 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if ip != 10 {
		t.Errorf("expected instruction pointer to be 10, but got %d", ip)
	}
}