package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

// 要求 log(n) 时间复杂度：意味着二分法（特殊的二分方式）
// 秘诀：不断查找“有序且符合区间”的数值，继用二分法判断
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 当数组只有两个元素时，特别处理
		if right-left == 1 {
			if nums[left] == target {
				return left
			} else if nums[right] == target {
				return right
			} else {
				return -1
			}
		}

		// 判断左边是否有序
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 否则，右边是有序的
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
