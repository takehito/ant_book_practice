package main

import "sort"

func sarumansArmy(r int, x []int) int {
	sort.IntSlice(x).Sort()
	var count, i int
	for i < len(x) {
		start := x[i]
		i++
		for i < len(x) && start+r < x[i] {
			i++
		}
		p := x[i-1]
		for i < len(x) && p+r >= x[i] {
			i++
		}
		count++
	}
	return count
}
