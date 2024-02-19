package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Original array:", arr)

	quickSort2(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:  ", arr)
}

// 递归调用自身进行排序
func quickSort2(arr []int, low, high int) {
	if low < high {
		// pivotIndex 为得到的基准元素索引
		pivotIndex := partition(arr, low, high)
		quickSort2(arr, low, pivotIndex-1)
		quickSort2(arr, pivotIndex+1, high)
	}
}

// 参数：切片、最低位、最高位
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
