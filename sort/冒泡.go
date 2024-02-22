package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 5, 4}
	bubbleSort(nums)
	for _, v := range nums {
		fmt.Print(v, "\t")
	}
}

func bubbleSort(nums []int) {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
