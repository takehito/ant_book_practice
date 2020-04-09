package fence

import "testing"

func TestFenceRepair(t *testing.T) {
	if got := fenceRepair([]int{8, 5, 8}); got != 34 {
		t.Errorf("%d\n", got)
	}
}
