package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k1, v := range nums {
		if k2, ok := m[target-v]; ok {
			return []int{k2, k1}
		}
		m[v] = k1
	}
	return []int{-1, -1}
}
