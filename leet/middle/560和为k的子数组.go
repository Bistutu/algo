package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 1, 3}, 4))
}

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := m[sum-k]; ok {
			count += m[sum-k]
		}
		m[sum]++
	}
	return count
}
