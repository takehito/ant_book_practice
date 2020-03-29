package main

type dfs struct {
	nums []int
	k int
}

func (d dfs) search(i int, sum int) bool {
	if i == len(d.nums) {
		return sum == d.k
	}

	if d.search(i+1, sum) {
		return true
	}

	if (d.search(i+1, sum+d.nums[i])) {
		return true
	}
	return false
}