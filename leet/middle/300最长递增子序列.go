package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	dp := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	maxSeries := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 如果  nums[i]>nums[j]，则
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > maxSeries {
					maxSeries = dp[i]
				}
			}
		}
	}
	return maxSeries
}
