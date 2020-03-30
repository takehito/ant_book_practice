package main

import "testing"

func TestCoin(t *testing.T) {
	coins := newCoins(3, 2, 1, 3, 0, 2)

	if got := coins.search(620); got != 6 {
		t.Errorf("%d\n", got)
	}
}
