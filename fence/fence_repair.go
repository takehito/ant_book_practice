package fence

func fenceRepair(l []int) int {
	var cost int

	for len(l) > 1 {
		mii1 := 0
		mii2 := 1
		if l[mii1] > l[mii2] {
			mii1, mii2 = mii2, mii1
		}

		for i := 2; i < len(l); i++ {
			if l[i] < l[mii1] {
				mii2 = mii2
				mii1 = i
			} else if l[i] < l[mii2] {
				mii2 = i
			}

		}

		t := l[mii1] + l[mii2]
		cost += t

		if mii1 == len(l)-1 {
			mii1, mii2 = mii2, mii1
		}
		l[mii1] = t
		l[mii2] = l[len(l)-1]
		l = l[:len(l)-1]
	}
	return cost
}
