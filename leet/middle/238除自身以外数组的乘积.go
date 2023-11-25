package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(num))
}

func productExceptSelf(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	result[0] = 1 // 初始化 [0] 为 1

	for i := 1; i < size; i++ {
		result[i] = nums[i-1] * result[i-1]
	}
	temp := 1
	for i := size - 1; i >= 0; i-- {
		result[i] = result[i] * temp
		temp *= nums[i]
	}
	return result
}
