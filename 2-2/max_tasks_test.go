package main

import "testing"

func TestMaxTasks(t *testing.T) {
	tasks := newTasks(
		[]task{
			task{start: 1, end: 3},
			task{start: 2, end: 5},
			task{start: 4, end: 7},
			task{start: 6, end: 9},
			task{start: 8, end: 10},
		},
	)

	if got := tasks.calc(); got != 3 {
		t.Errorf("%d\n", got)
	}
}
