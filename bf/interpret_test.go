package bf

import (
	"testing"
	"time"
)

func TestInterpret(t *testing.T) {
	stopChan := make(chan(struct{}))
	timeout := time.Second // Unused when Interpret is not called from WrappedInterpretWithTimeout. Just supplied here to match the signature.

	s, err := Interpret("+>++++++++[<++++++++>-]<.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "A" {
		t.Errorf("Expected 'A', but got '%q'", s)
	}

	s, err = Interpret("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'AAA', but got %s", err.Error())
	}
	if s != "AAA" {
		t.Errorf("Expected 'AAA', but got '%q'", s)
	}

	s, err = Interpret("+>++++++++[<++++++++>-]+++[<.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'AAA', but got %s", err.Error())
	}
	if s != "AAA" {
		t.Errorf("Expected 'AAA', got '%q'", s)
	}

	s, err = Interpret("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'ABC', but got %s", err.Error())
	}
	if s != "ABC" {
		t.Errorf("Expected 'ABC', but got '%q'", s)
	}

	s, err = Interpret(">++++++++[<++++++++>-]+++[<+.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'ABC', but got %s", err.Error())
	}
	if s != "ABC" {
		t.Errorf("Expected 'ABC', but got '%q'", s)
	}

	s, err = Interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected ''Hello World!\\n'', but got %s", err.Error())
	}
	if s != "Hello World!\n" {
		t.Errorf("Expected 'Hello World!\\n', but got '%q'", s)
	}

	s, err = Interpret(">+++++++[<+++++++>-]+++++++++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected '123456789', but got %s", err.Error())
	}
	if s != "123456789" {
		t.Errorf("Expected '123456789', got '%q'", s)
	}

	s, err = Interpret(">+++++++<->[<+++++++>-]++++++++++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected '0123456789', but got %s", err.Error())
	}
	if s != "0123456789" {
		t.Errorf("Expected '0123456789', got '%q'", s)
	}

	s, err = Interpret("++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'I LOVE YOU', but got %s", err.Error())
	}
	if s != "I LOVE YOU" {
		t.Errorf("Expected 'I LOVE YOU', got '%q'", s)
	}

	// Test array is at least 30_000 bytes long.
	s, err = Interpret("++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected '#\\n', but got %s", err.Error())
	}
	if s != "#\n" {
		t.Errorf("Expected '#\\n', but got '%q'", s)
	}

	s, err = Interpret("++>+++++[<+>-]++++++++[<++++++>-]<.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected '7', but got %s", err.Error())
	}
	if s != "7" {
		t.Errorf("Expected '7', but got '%q'", s)
	}
}
