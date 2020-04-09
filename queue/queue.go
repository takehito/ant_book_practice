package main

import (
	"fmt"
)

type queue struct {
	size     int
	value    []int
	location int
}

func newQueue(n int) *queue {
	return &queue{
		size:     n,
		value:    make([]int, n),
		location: 0,
	}
}

func (q *queue) push(v int) {
	if q.location == q.size {
		return
	}

	q.value[q.location] = v

	q.location++
}

func (q *queue) pop() {
	q.value = q.value[1:]
	fmt.Println(q.value)
	q.location--
}

func (q *queue) front() int {
	if q.location == 0 {
		return 0
	}

	return q.value[0]
}
