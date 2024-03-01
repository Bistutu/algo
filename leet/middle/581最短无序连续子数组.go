package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	left, right := 0, len(nums)-1
	for nums[left] == temp[left] {
		left++
	}
	for nums[right] == temp[right] {
		right--
	}
	return right - left + 1
}
