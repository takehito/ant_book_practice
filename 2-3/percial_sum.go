package main

type am struct {
	a int
	m int
}

type ams []am

func (n ams) percialSum(want int) bool {
	dp := make([][]bool, len(n)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, want+1)
	}
	dp[0][0] = true

	for i := 0; i < len(n); i++ {
		for j := 0; j <= want; j++ {
			for k := 0; k <= n[i].m && k*n[i].a <= j; k++ {
				dp[i+1][j] = dp[i+1][j] || dp[i][j-k*n[i].a]
			}
		}
	}

	return dp[len(n)][want]
}
