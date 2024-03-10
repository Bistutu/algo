package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

// 秘诀：先乘左边、再乘右边
func productExceptSelf(nums []int) []int {
	rs := make([]int, len(nums))
	rs[0] = 1 // 初始化
	for i := 1; i < len(nums); i++ {
		rs[i] = rs[i-1] * nums[i-1]
	}

	temp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 rs[i]，右边的乘积为 temp
		rs[i] = rs[i] * temp
		temp *= nums[i]
	}
	return rs
}
