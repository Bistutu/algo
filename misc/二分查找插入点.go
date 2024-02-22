package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, 6, 7}
	fmt.Println(binarySearchDot(nums, 4))
}

// 二分查找
func binarySearchDot(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
