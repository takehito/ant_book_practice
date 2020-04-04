package main

import "math"

type napsack struct {
	weight int
	value  int
}

func maxValue(n []napsack, memos []napsack, i int, maxWeight int) int {
	if i < len(memos) && memos[i].weight == maxWeight {
		return memos[i].value
	}
	var result int

	if i == len(n) {
		// もう品物は残っていない
		result = 0
	} else if maxWeight < n[i].weight {
		// この品物は入らない
		result = maxValue(n, memos, i+1, maxWeight)
	} else {
		result = int(math.Max(float64(maxValue(n, memos, i+1, maxWeight)), float64(maxValue(n, memos, i+1, maxWeight-n[i].weight)+n[i].value)))
	}

	memos[i].value = result
	memos[i].weight = maxWeight
	return result
}

type napsacks []napsack

func (n napsacks) unlimitedNapsack() int {
	dp := make([][]int, len(n)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(n)+1)
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			for k := 0; k*n[i].weight <= j; k++ {
				dp[i+1][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-k*n[i].weight]+n[i].value)))
			}
		}
	}

	return dp[len(n)][len(n)]
}
