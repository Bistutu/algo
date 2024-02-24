package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

// 秘诀：套 set、找起点、判长度
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}
	maxLong := 0
	for num := range m { // 这里 range 的应为 m
		// 如果是起点（即前面 num-1 没有元素）
		if _, ok := m[num-1]; !ok {
			length := 0
			current := num
			for {
				if _, ok := m[current]; ok {
					length++
					current++
				} else {
					break
				}
			}
			if length > maxLong {
				maxLong = length
			}
		}

	}
	return maxLong
}
