package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1}
	quickSort(nums, 0, len(nums)-1)
	for _, v := range nums {
		fmt.Printf("%d\t", v)
	}
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := handler(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func handler(nums []int, left, right int) int {
	// 这里基准元素默认为 nums[left]
	i, j := left, right
	for i < j {
		// 从左至右，查找首个小于基准数的元素
		for i < j && nums[j] >= nums[left] {
			j--
		}
		// 从右至左，查找首个大于基准数的元素
		for i < j && nums[i] <= nums[left] {
			i++
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	// 返回基准元素的索引
	return i
}
