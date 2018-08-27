package ds

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	v := s.Pop().(int)
	if v != 3 {
		t.Errorf("Failed. Expected: 3, Got: %v", v)
	}
	if s.Len() != 2 {
		t.Errorf("Failed. Expected: 2, Got: %v", s.Len())
	}
	s.Pop()
	s.Pop()
	i := s.Pop()
	if i != nil {
		t.Errorf("Failed. Expected: nil, Got: %v", v)
	}
}

func TestIsValidParentheses(t *testing.T) {
	s := "()[]{}"
	if !IsValidParentheses(s) {
		t.Error("Failed")
	}

	s = "([)]"
	if IsValidParentheses(s) {
		t.Error("Failed")
	}
}

func TestAddString(t *testing.T) {
	s := AddString("123", "985")
	if s != "1108" {
		t.Errorf("Failed. Expected: 168. Got: %s", s)
	}
}
