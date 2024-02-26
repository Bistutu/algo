package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3

	result := searchMatrix(matrix, target)
	fmt.Printf("The matrix contains %d: %v\n", target, result)

	target = 13
	result = searchMatrix(matrix, target)
	fmt.Printf("The matrix contains %d: %v\n", target, result)
}

// Z 字形查找效率最高
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1        // 获取右上角坐标
	for x < m && y >= 0 { // 注意判断条件
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
