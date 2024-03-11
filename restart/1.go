package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, k}
		}

		m[v] = k
	}
	return nil
}
