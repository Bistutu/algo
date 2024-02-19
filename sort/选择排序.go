package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 5, 4}
	selectSort(nums)
	for k, v := range nums {
		fmt.Println(k, v)
	}
}

func selectSort(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	var min int // 记录最小值的下标
	// 循环，每一轮都把最小的放到最前面
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
