package main

import "testing"

func TestCommonString(t *testing.T) {
	if got := commonString("abcd", "becd"); got != 3 {
		t.Errorf("%d\n", got)
	}
}
