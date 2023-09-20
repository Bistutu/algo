package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 5, 4}
	insertSort(nums)
	for _, v := range nums {
		fmt.Printf("%d\t", v)
	}
}

// 插入排序，最小的排在最前面
func insertSort(nums []int) {
	// 外循环，[1,end]
	for i := 1; i < len(nums); i++ {
		base := nums[i] // 基准元素，默认为 nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}
