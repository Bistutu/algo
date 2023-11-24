package main

import "fmt"

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}
func main() {
	// 示例数组
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)  // 排序
	fmt.Println(nums) // 输出排序后的数组：[0, 0, 1, 1, 2, 2]

	nums = []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums) // 输出排序后的数组：[0, 1, 2]
}
