package main

import "testing"

func TestDeepSearch(t *testing.T) {
	d1 := dfs{
		nums: []int{1, 2, 4, 7},
		k:    13,
	}
	if got := d1.search(0, 0); got != true {
		t.Errorf("1 miss")
	}
	d1.k = 15
	if got := d1.search(0, 0); got != false {
		t.Errorf("2 miss")
	}
}
