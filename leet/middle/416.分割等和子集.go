package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5})) // true
	fmt.Println(canPartition([]int{1, 2, 3, 5}))  // false
}

func canPartition(nums []int) bool {
	// 如果和是奇数，则不能分成 2 个相等的子集
	count := 0
	for _, v := range nums {
		count += v
	}
	if count%2 != 0 {
		return false
	}
	target := count / 2 // 最终分成的两部分后，均分的值
	size := len(nums)

	dp := make([][]bool, size+1)
	for i := 0; i <= size; i++ {
		dp[i] = make([]bool, target+1)
	}

	dp[0][0] = true

	for i := 1; i <= size; i++ {
		num := nums[i-1]
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if num <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-num]
			}
		}
	}

	return dp[size][target]
}
