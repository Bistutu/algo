package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 5, 6, 7}
	fmt.Println(binarySearchRightEdge(nums, 3))
}

// 二分查找左边界
func binarySearchLeftEdge(nums []int, target int) int {
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
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 右边界
func binarySearchRightEdge(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
