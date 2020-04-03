package main

import "testing"

func TestNapsack(t *testing.T) {
	n := []napsack{
		napsack{weight: 2, value: 3},
		napsack{weight: 1, value: 2},
		napsack{weight: 3, value: 4},
		napsack{weight: 2, value: 2},
	}
	memos := make([]napsack, len(n)+1)
	if got := maxValue(n, memos, 0, 5); got != 7 {
		t.Errorf("%d\n", got)
	}
}
