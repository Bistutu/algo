package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// 允许 3 种操作：插入、删除、替换
func minDistance(s string, t string) int {
	sSize := len(s) + 1
	tSize := len(t) + 1
	dp := make([][]int, sSize)
	for i := 0; i < sSize; i++ {
		dp[i] = make([]int, tSize)
	}
	for i := 0; i < sSize; i++ {
		dp[i][0] = i
	}
	for i := 0; i < tSize; i++ {
		dp[0][i] = i
	}

	for i := 1; i < sSize; i++ {
		for j := 1; j < tSize; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = mins(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[sSize-1][tSize-1]
}

func mins(a, b, c int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}
