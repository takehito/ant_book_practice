package main

import "testing"

func TestPercialSum(t *testing.T) {
	n := []am{
		am{a: 3, m: 3},
		am{a: 5, m: 2},
		am{a: 8, m: 2},
	}
	if got := ams(n).percialSum(17); got != true {
		t.Error(got)
	}
}

func TestNerPercialSum(t *testing.T) {
	n := []am{
		am{a: 3, m: 3},
		am{a: 5, m: 2},
		am{a: 8, m: 2},
	}
	if got := ams(n).newPercialSum(17); got != true {
		t.Error(got)
	}
}
