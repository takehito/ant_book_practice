package main

type stack struct {
	previous *stack
	current  *stack
	value    int
}

func (s *stack) push(n int) {
	s.previous = s.current
	s.current = &stack{value: n, previous: s.current}
}

func (s *stack) pop() {
	s.current = s.previous
	s.previous = s.previous.previous
}

func (s *stack) top() int {
	return s.current.value
}
