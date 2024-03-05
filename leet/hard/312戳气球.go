package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8})) // 167
}

func maxCoins(nums []int) int {
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
