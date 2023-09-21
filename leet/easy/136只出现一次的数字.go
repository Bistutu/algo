package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2, 5, 5}))
}

// 异或，0 ^ a ^ b ^ b = a
func singleNumber(nums []int) int {
	result := 0
	count := len(nums)
	for i := 0; i < count; i++ {
		result ^= nums[i]
	}
	return result
}
