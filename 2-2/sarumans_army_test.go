package main

import "testing"

func TestSarumansArmy(t *testing.T) {
	if got := sarumansArmy(10, []int{1, 7, 15, 20, 30, 50}); got != 3 {
		t.Errorf("%d\n", got)
	}
}
