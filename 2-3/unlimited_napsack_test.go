package main

import "testing"

func TestUnlimitedNapsack(t *testing.T) {
	n := []napsack{
		napsack{weight: 3, value: 4},
		napsack{weight: 4, value: 5},
		napsack{weight: 2, value: 3},
	}
	if got := napsacks(n).unlimitedNapsack(); got != 10 {
		t.Errorf("%d\n", got)
	}
}
