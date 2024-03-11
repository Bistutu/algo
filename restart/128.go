package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	maxLength := 0
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		// 如果是从头开始的元素
		if !m[v-1] {
			count := 0
			for i := v; m[i]; i++ {
				count++
			}
			maxLength = max(maxLength, count)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
