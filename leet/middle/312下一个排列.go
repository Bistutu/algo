package main

import "fmt"

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{3, 2, 1})
}

// 在函数里输出
func nextPermutation(nums []int) {
	// 秘诀：基于一个数学基础
	// 1、从右找到第一个 nums[i] < nums[i+1] 的元素
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 2、如果不是完全的降序
	if i >= 0 {
		j := len(nums) - 1
		// 3、从右至左找到第一个比 nums[i] 大的元素
		for nums[i] >= nums[j] {
			j--
		}
		// 4、交换这两个元素
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 5、反转 i 后的所有元素
	reverse(nums, i+1, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
