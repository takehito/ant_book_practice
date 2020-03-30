package main

import "math"

type coin struct {
	value int
	num   int
}

type coins []coin

func newCoins(i, j, k, l, m, n int) coins {
	return []coin{
		coin{value: 1, num: i},
		coin{value: 5, num: j},
		coin{value: 10, num: k},
		coin{value: 50, num: l},
		coin{value: 100, num: m},
		coin{value: 500, num: n},
	}
}

func (c coins) search(value int) int {
	var count int
	for i := len(c) - 1; i >= 0; i-- {
		// 枚数を求める
		t := math.Min(float64(value/c[i].value), float64(c[i].num))
		value -= c[i].value * int(t)
		count += int(t)
	}
	return count
}
