package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}

func searchRange(nums []int, target int) []int {
	start := sort.SearchInts(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.SearchInts(nums[start:], target) + 1 + start
	if end == len(nums[start:]) || nums[end] != target {
		end = start
	}

	return []int{start, end}
}
