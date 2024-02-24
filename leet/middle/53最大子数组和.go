package main

import "fmt"

func main() {
	// 每个数字都会有以它结尾的情况，所以这里只不过是换了种思想（即使用了这种想法）。
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	dp := make(map[int]int, len(nums))
	dp[0] = nums[0]
	maxNum := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxNum = max(maxNum, dp[i])
	}
	return maxNum
}
