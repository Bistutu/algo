package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 num1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// m、n 分别为两个数组的长度
	m, n := len(nums1), len(nums2)
	// imin、imax 用于在 num1 上进行二分查找，halfLen 为两个数组合并后左半部分的长度
	imin, imax, halfLen := 0, m, (m+n+1)/2

	// 核心编程：二分查找
	for imin < imax {
		i := (imin + imax) / 2 // i 是 num1 的分割点

	}

	return 0
}

// 暴力解法，目标 O(n)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	nums := make([]int, 0, l1+l2)
	for i < l1 || j < l2 {
		if i < l1 && j == l2 {
			nums = append(nums, nums1[i])
			i++
		} else if i == l1 && j < l2 {
			nums = append(nums, nums2[j])
			j++
		} else {
			if nums1[i] < nums2[j] {
				nums = append(nums, nums1[i])
				i++
			} else {
				nums = append(nums, nums2[j])
				j++
			}
		}
	}
	left, right := 0, 0
	if len(nums)%2 == 0 {
		left = len(nums)/2 - 1
		right = len(nums) / 2
	} else {
		left, right = len(nums)/2, len(nums)/2
	}
	return float64(nums[left]+nums[right]) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
