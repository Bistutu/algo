package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(MergeSort(arr))
}

// MergeSort 对一个整数切片执行归并排序
func MergeSort(arr []int) []int {
	// 基础情况：一个长度为零或一的列表默认是已排序的。
	if len(arr) <= 1 {
		return arr
	}

	// 将未排序的列表分为两个大致相等的部分。
	mid := len(arr) / 2
	leftHalf := arr[:mid]
	rightHalf := arr[mid:]

	// 递归地对两个子列表进行排序。
	leftSorted := MergeSort(leftHalf)
	rightSorted := MergeSort(rightHalf)

	// 然后合并两个现在已排序的子列表。
	return merge(leftSorted, rightSorted)
}

// merge 接受两个已排序的切片，并返回一个新的已排序切片，其中包含两个输入切片的所有元素。
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果 'left' 或 'right' 切片中还有剩余的元素，将它们追加到 'result' 中。
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
