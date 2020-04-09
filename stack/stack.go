package stack

type stack struct {
	value []int
}

func (s *stack) push(n int) {
	s.value = append(s.value, n)
}

func (s *stack) pop() {
	s.value = s.value[:len(s.value)-1]
}

func (s *stack) top() int {
	return s.value[len(s.value)-1]
}
