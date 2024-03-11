package main

import (
	"fmt"
)

func main() {
	//fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if val, ok := m[v]; ok && abs(val, i) <= k {
			return true
		}
		m[v] = i
	}
	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
