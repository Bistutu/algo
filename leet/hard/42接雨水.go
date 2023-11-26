package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	// 双指针
	left, right := 0, len(height)-1
	count := 0
	maxLeft, maxRight := 0, 0

	for left < right {
		if height[left] < height[right] {
			if maxLeft > height[left] {
				count += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
			left++
		} else {
			if maxRight > height[right] {
				count += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
			right--
		}
	}
	return count
}
