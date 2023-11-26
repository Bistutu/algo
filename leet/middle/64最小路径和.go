package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	endx, endy := len(grid)-1, len(grid[0])-1
	mem := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		mem[i] = make([]int, len(grid[0]))
	}
	return minPathSumDFS(grid, mem, endx, endy)
}

func minPathSumDFS(grid, mem [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	if mem[i][j] != 0 {
		return mem[i][j]
	}
	left := minPathSumDFS(grid, mem, i-1, j)
	top := minPathSumDFS(grid, mem, i, j-1)

	score := min(left, top) + grid[i][j]
	mem[i][j] = score
	return score
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
