package main

// 1 + 2 + 3 + ... + n
func beisen2(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}

// n + ... + 1
// n + beisen3(n-1)
func beisen3(n int) int {
	if n == 1 {
		return 1
	}
	return n + beisen3(n-1)
}
