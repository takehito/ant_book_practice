package main

import (
	"math"
)

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

func (n napsacks) initDP(maxWeight int) [][]int {
	dp := make([][]int, len(n)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	return dp
}

func (n napsacks) unlimitedNapsack(maxWeight int) int {
	dp := n.initDP(maxWeight)

	for i := 0; i < len(n); i++ {
		for j := 0; j <= maxWeight; j++ {
			if j < n[i].weight {
				dp[i+1][j] = dp[i][j]
			} else {
				a := dp[i][j]
				b := dp[i+1][j-n[i].weight] + n[i].value
				dp[i+1][j] = int(math.Max(float64(a), float64(b)))
			}
		}
	}

	return dp[len(n)][maxWeight]
}

func (n napsacks) napsack01(maxWeight int) int {
	dp := make([]int, maxWeight+1)
	for i := 0; i < len(n); i++ {
		for j := maxWeight; j >= n[i].weight; j-- {
			update := dp[j-n[i].weight] + n[i].value
			dp[j] = int(math.Max(float64(dp[j]), float64(update)))
		}
	}
	return dp[maxWeight]
}

func (n napsacks) unlimitedNapsackFixed(maxWeight int) int {
	dp := make([]int, maxWeight+1)
	for i := 0; i < len(n); i++ {
		for j := n[i].weight; j <= maxWeight; j++ {
			update := dp[j-n[i].weight] + n[i].value
			dp[j] = int(math.Max(float64(dp[j]), float64(update)))
		}
	}
	return dp[maxWeight]
}
