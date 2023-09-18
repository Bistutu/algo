package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 5, 4}
	bubbleSort(nums)
	for k, v := range nums {
		fmt.Println(k, v)
	}

}

func bubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
