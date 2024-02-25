package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0 // 此处应该抛出错误
	}

	sort.Ints(nums)
	if len(nums) < 3 {
		return nums[len(nums)-1]
	}
	times := 1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			times++
		}
		if times == 3 {
			return nums[i-1]
		}
	}
	// 最终情景距离：[1,2,2,2,2]，则返回最后的数字 nums[len(nums)-1] 即 2
	return nums[len(nums)-1]
}
