package algo

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
