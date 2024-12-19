package bf

import (
	"errors"
	"testing"
	"time"
)

func TestInterpret(t *testing.T) {
	stopChan := make(chan(struct{}))
	timeout := time.Second // Unused when Interpret is not called from WrappedInterpretWithTimeout. Just supplied here to match the signature.

	tests := []struct {
		name        string
		input       string
		expectedStr string
		expectedErr error
	}{
		{"Hello World", "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", "Hello World!\n", nil},
		{"A", "+>++++++++[<++++++++>-]<.", "A", nil},
		{"Long AAA", "+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.>-]", "AAA", nil},
		{"Short AAA", "+>++++++++[<++++++++>-]+++[<.>-]", "AAA", nil},
		{"Long ABC", "+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.+>-]", "ABC", nil},
		{"Short ABC", ">++++++++[<++++++++>-]+++[<+.>-]", "ABC", nil},
		{"123456789", ">+++++++[<+++++++>-]+++++++++[<.+>-]", "123456789", nil},
		{"0123456789", ">+++++++<->[<+++++++>-]++++++++++[<.+>-]", "0123456789", nil},
		{"I LOVE YOU", "++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++.", "I LOVE YOU", nil},
		{"Check array is at least 30_000 bytes long", "++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.", "#\n", nil},
		{"2 + 5 = 7", "++>+++++[<+>-]++++++++[<++++++>-]<.", "7", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Interpret(tt.input, stopChan, timeout)

			if output != tt.expectedStr {
				t.Errorf("Interpret(%q, , ) string = %q, want %q", tt.input, output, tt.expectedStr)
			}

			if (err != nil || tt.expectedErr != nil) && !errors.Is(err, tt.expectedErr) {
				t.Errorf("Interpret(%q, , ) error = %q, want %q", tt.input, err, tt.expectedErr)
			}
		})
	}
}
