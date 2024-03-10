package main

import "fmt"

func main() {
	//matrix := [][]byte{
	//	{'1', '0', '1', '0', '0'},
	//	{'1', '0', '1', '1', '1'},
	//	{'1', '1', '1', '1', '1'},
	//	{'1', '0', '0', '1', '0'},
	//}
	matrix := [][]byte{{'1', '0'}}
	result := maximalRectangle(matrix)
	fmt.Println("最大矩形面积是:", result)
}

// 暴力解法
func maximalRectangle2(matrix [][]byte) int {
	return 0
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	height := make([]int, len(matrix[0])) // “构建”柱状图的 height
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// 计算当前形成的最大矩形面积，并更新 maxArea
		maxArea = max(maxArea, largestRectangleArea(height))
	}

	return maxArea
}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
