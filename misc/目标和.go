package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
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
		t := nums[0]
		nums = nums[1:]
		dfs(nums, cur+t)
		dfs(nums, cur-t)
	}
	dfs(nums, 0)
	return count
}
