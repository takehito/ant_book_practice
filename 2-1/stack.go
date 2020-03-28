package main

type stack struct {
	size     int
	value    []int
	location int
}

func newStack(n int) *stack {
	return &stack{
		size:  n,
		value: make([]int, n),
	}
}

func (s *stack) push(n int) {
	if s.location == s.size-1 {
		return
	}
	s.value[s.location] = n
	s.location++
}

func (s *stack) pop() {
	if s.location == 0 {
		return
	}
	s.location--
}

func (s *stack) top() int {
	return s.value[s.location]
}
