package napsack

import "testing"

func TestUnlimitedNapsack(t *testing.T) {
	n := []napsack{
		napsack{weight: 3, value: 4},
		napsack{weight: 4, value: 5},
		napsack{weight: 2, value: 3},
	}
	if got := napsacks(n).unlimitedNapsack(7); got != 10 {
		t.Errorf("%d\n", got)
	}
}

func TestUnlimitedNapsackFixed(t *testing.T) {
	n := []napsack{
		napsack{weight: 3, value: 4},
		napsack{weight: 4, value: 5},
		napsack{weight: 2, value: 3},
	}
	if got := napsacks(n).unlimitedNapsackFixed(7); got != 10 {
		t.Errorf("%d\n", got)
	}
}

func TestNapsack01(t *testing.T) {
	n := []napsack{
		napsack{weight: 2, value: 3},
		napsack{weight: 1, value: 2},
		napsack{weight: 3, value: 4},
		napsack{weight: 2, value: 2},
	}
	if got := napsacks(n).napsack01(5); got != 7 {
		t.Errorf("%d\n", got)
	}
}

func TestMaxValue(t *testing.T) {
	n := []napsack{
		napsack{weight: 2, value: 3},
		napsack{weight: 1, value: 2},
		napsack{weight: 3, value: 4},
		napsack{weight: 2, value: 2},
	}
	if got := napsacks(n).getMaxValue(); got != 4 {
		t.Errorf("%d\n", got)
	}
}

func TestNapsack02(t *testing.T) {
	n := []napsack{
		napsack{weight: 2, value: 3},
		napsack{weight: 1, value: 2},
		napsack{weight: 3, value: 4},
		napsack{weight: 2, value: 2},
	}
	if got := napsacks(n).napsack02(5); got != 7 {
		t.Errorf("%d\n", got)
	}
}
