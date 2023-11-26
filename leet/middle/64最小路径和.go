package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
func minPathSum(grid [][]int) int {
	y := len(grid)
	x := len(grid[0])

	dp := make([][]int, y)
	for i := 0; i < y; i++ {
		dp[i] = make([]int, x)
	}
	dp[0][0] = grid[0][0]

	// 计算首行
	for j := 1; j < x; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 计算首列
	for i := 1; i < y; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 计算其余单元格
	for i := 1; i < y; i++ {
		for j := 1; j < x; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[y-1][x-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
