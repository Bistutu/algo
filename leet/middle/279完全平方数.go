package main

import "fmt"

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 最差的情况
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
