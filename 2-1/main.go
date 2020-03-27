package main

func fullSearchFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	result := fullSearchFibonacci(n-1) + fullSearchFibonacci(n-2)
	return result
}

func main() {
	fullSearchFibonacci(6)
}
