package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

//
//func maxArea(height []int) int {
//	maxArea := 0
//	for i := 0; i < len(height)-1; i++ {
//		for j := i + 1; j < len(height); j++ {
//			if (j-i)*min(height[i], height[j]) > maxArea {
//				maxArea = (j - i) * min(height[i], height[j])
//			}
//		}
//	}
//	return maxArea
//}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
