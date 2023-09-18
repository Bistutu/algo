package main

import "fmt"

func main() {
	fmt.Println(BinarySearch([]int{10, 20, 30, 40, 60, 70}, 0))
}

func BinarySearch(num []int, target int) int {
	if num == nil {
		return -1
	}
	//
	var i, j = 0, len(num) - 1
	var middle int
	for i <= j {
		// 求中点（算式防止直接 j+i 导致的int越界）
		middle = i + (j-i)/2
		if target == num[middle] {
			return middle
		} else if target > num[middle] {
			i = middle + 1
		} else {
			j = middle - 1
		}
	}
	return -1
}
