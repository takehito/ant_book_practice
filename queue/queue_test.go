package main

import "testing"

func TestQueue(t *testing.T) {
	q := newQueue(3)
	q.push(1)
	q.push(2)
	q.push(3)
	if got := q.front(); got != 1 {
		t.Errorf("%d\n", got)
	}
	q.pop()
	if got := q.front(); got != 2 {
		t.Errorf("%d\n", got)
	}
	q.pop()
	if got := q.front(); got != 3 {
		t.Errorf("%d\n", got)
	}
}
