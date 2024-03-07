package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	count := 0 // 记录所接雨水总量
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0 // 左右最高柱子的高度

	for left < right {
		// 如果左边的柱子高度小于右边的柱子高度
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				count += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				count += maxRight - height[right]
			}
			right--
		}
	}

	return count
}
