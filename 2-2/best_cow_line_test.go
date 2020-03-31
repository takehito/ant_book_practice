package main

import "testing"

func TestBestCowLine(t *testing.T) {
	s := "ACDBCB"
	if got := BestCowLine(s); got != "ABCBCD" {
		t.Errorf("%s\n", got)
	}
}
