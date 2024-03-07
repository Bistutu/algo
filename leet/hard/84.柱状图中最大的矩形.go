package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 输出: 10
	fmt.Println(largestRectangleArea([]int{2, 4}))             // 输出: 4
	fmt.Println(largestRectangleArea([]int{1, 1}))             // 输出: 2
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 5})) // 输出: 6
}

// 秘诀：使用一个「单调栈」存储下标
/*第一轮和第二轮的 for 循环分别处理不同的场景
1、第一轮 for 循环：处理大多数的情况，包括计算以每根柱子为高的矩形区域的面积。
2、第二轮 for 循环：处理剩余的情况，即在柱状图的末端仍在栈中的柱子。
*/

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]int, 0)

	// 第一轮 for 循环
	for i := 0; i < len(heights); i++ {
		// 如果栈顶元素大于 height[i]，则说明找到了「每根柱子为高的矩形区域的面积」的情况
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			// 弹出元素
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, heights[temp]*width)
		}
		stack = append(stack, i)
	}

	// 第二轮 for 循环，此时单调栈建立完成，且已经获得全部「每根柱子为高的矩形区域的面积」的场景，接下来考虑其他场景
	// （这次相当于是从右前左，向前处理）
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := len(heights) // 初始宽度为柱状图的高度

		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		maxArea = max(maxArea, heights[temp]*width)
	}

	return maxArea
}

// 暴力解法（超出时间限制）
func largestRectangleArea2(heights []int) int {
	size := len(heights)
	maxArea := 0
	for i := 0; i < size; i++ {
		left := i
		right := i + 1

		for j := i; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			left = j
		}

		for j := i; j < size; j++ {
			if heights[j] < heights[i] {
				break
			}
			right = j + 1
		}

		area := heights[i] * (right - left)
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
