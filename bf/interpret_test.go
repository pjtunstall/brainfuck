package bf

import (
	"testing"
	"time"
)

func TestInterpret(t *testing.T) {
	stopChan := make(chan(struct{}))
	timeout := time.Second // Unused. Just supplied to match the signature.

	s, err := Interpret("+>++++++++[<++++++++>-]<.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "A" {
		t.Errorf("Expected 'A', but got '%q'", s)
	}

	s, err = Interpret("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "AAA" {
		t.Errorf("Expected 'AAA', but got '%q'", s)
	}

	s, err = Interpret("+>++++++++[<++++++++>-]+++[<.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "AAA" {
		t.Errorf("Expected 'AAA', got '%q'", s)
	}

	s, err = Interpret("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>+++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "ABC" {
		t.Errorf("Expected 'ABC', but got '%q'", s)
	}

	s, err = Interpret(">++++++++[<++++++++>-]+++[<+.>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "ABC" {
		t.Errorf("Expected 'ABC', but got '%q'", s)
	}

	s, err = Interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "Hello World!\n" {
		t.Errorf("Expected 'Hello World!\n', but got '%q'", s)
	}

	s, err = Interpret(">+++++++[<+++++++>-]+++++++++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "123456789" {
		t.Errorf("Expected '123456789', got '%q'", s)
	}

	s, err = Interpret(">+++++++<->[<+++++++>-]++++++++++[<.+>-]", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "0123456789" {
		t.Errorf("Expected '0123456789', got '%q'", s)
	}

	s, err = Interpret("++++++++[>+++++++++>++++++++>++++<<<-]>+.>>.<<+++.+++.+++++++.>+++++.>.<<+++.----------.++++++.", stopChan, timeout)
	if err != nil {
		t.Errorf("Expected 'A', but got %s", err.Error())
	}
	if s != "I LOVE YOU" {
		t.Errorf("Expected 'I LOVE YOU', got '%q'", s)
	}
}