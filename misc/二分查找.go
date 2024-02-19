package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(nums, 7))
}

// 二分查找
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
