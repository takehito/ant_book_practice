package sequence

import "math"

// MaxLength is returning Longest increasing subsequence
func MaxLength(n []int) int {
	dp := make([]int, len(n))
	var res int
	for i := 0; i < len(n); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if n[j] < n[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}

// NewMaxLength is returning Longest increasing subsequence
func NewMaxLength(n []int) int {
	const INF = 10000000
	dp := make([]int, len(n))
	for i := 0; i < len(n); i++ {
		dp[i] = INF
	}

	for i := 0; i < len(n); i++ {
		lower_bound(dp, dp+n, n[i]) = n[i]
	}
	return lower_bound(dp, dp+n, INF)
}
