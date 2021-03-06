package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var s stack
	s.push(1)
	s.push(2)
	s.push(3)
	if got := s.top(); got != 3 {
		t.Errorf("got: %d, expected: 3\n", got)
	}
	s.pop()
	if got := s.top(); got != 2 {
		t.Errorf("got: %d, expected: 2\n", got)
	}
	s.pop()
	if got := s.top(); got != 1 {
		t.Errorf("got: %d, expected: 1\n", got)
	}
}
