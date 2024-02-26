package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{0, 2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	low, high, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 当  nums[i] < 0，交换
		if nums[i] < 0 {
			low, high = high, low
		}
		high = max(nums[i], high*nums[i])
		low = min(nums[i], low*nums[i])

		result = max(result, high)
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
