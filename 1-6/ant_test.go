package ant

import "testing"

func TestAntWalkMin(t *testing.T) {
	if got := antWalkMin(10, []int{2, 6, 7}); got != 5 {
		t.Error(got)
	}
}

func TestAntWalkMax(t *testing.T) {
	if got := antWalkMax(10, []int{2, 6, 7}); got != 9 {
		t.Error(got)
	}
}
