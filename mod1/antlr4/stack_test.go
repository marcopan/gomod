package antlr4

import "testing"

func TestNewStack(t *testing.T) {
	newStack := NewStack()
	newStack.Push("test1")
	pop := newStack.Pop()
	if s, ok := pop.(string); ok {
		if s != "test" {
			t.Errorf("stack Pop error %s", s)
		}
	}

}
