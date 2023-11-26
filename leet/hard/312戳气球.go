package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8})) // 167
}

func maxCoins(nums []int) int {
	size := len(nums)
	dp := make([][]int, size+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, size+1)
	}
	// dp 表初始化完毕
	// TODO
	return 0
}

func getScore(nums []int, i int) int {
	left, right := 1, 1
	if i-1 > 0 {
		left = nums[i-1]
	}
	if i+1 < len(nums) {
		right = nums[i+1]
	}
	return left * nums[i] * right
}
