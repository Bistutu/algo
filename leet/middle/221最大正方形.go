package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix)) // 输出: 4
}

// 秘诀：判断 3 个点位，并取最小值（因为依靠该值才能判断组不组的成正方向）
func maximalSquare(matrix [][]byte) int {
	maxLength := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// 当为 '1' 时才统计
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = mins(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				}
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}
	return maxLength * maxLength
}

func mins(a, b, c int) int {
	minNum := a
	if b < minNum {
		minNum = b
	}
	if c < minNum {
		minNum = c
	}
	return minNum
}
