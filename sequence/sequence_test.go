package sequence

import "testing"

func TestMaxLength(t *testing.T) {
	if got := MaxLength([]int{4, 2, 3, 1, 5}); got != 3 {
		t.Error(got)
	}
}

func TestNewMaxLength(t *testing.T) {
	if got := NewMaxLength([]int{4, 2, 3, 1, 5}); got != 3 {
		t.Error(got)
	}
}
