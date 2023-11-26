package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(nums, 3)) // 输出 5，表示有 5 种方式
}

func findTargetSumWays(nums []int, target int) int {
	count := 0
	var dfs func(nums []int, cur int)
	dfs = func(nums []int, cur int) {
		if len(nums) == 0 {
			if cur == target {
				count++
			}
			return
		}
		temp := nums[0]
		nums = nums[1:]
		dfs(nums, cur+temp)
		dfs(nums, cur-temp)
	}
	dfs(nums, 0)
	return count
}
