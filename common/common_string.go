package common

import (
	"fmt"
	"math"
)

func commonString(s string, t string) int {
	dp := make([][]int, len(s)+1)

	fmt.Printf("%d\n", len(dp))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
		fmt.Println(len(dp[i]))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = int(math.Max(float64(dp[i][j+1]), float64(dp[i+1][j])))
			}
		}
	}
	return dp[len(s)][len(t)]
}
