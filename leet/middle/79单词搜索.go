package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED"))
}

// exist 函数用于判断二维字符数组 board 中是否存在字符串 word
func exist(board [][]byte, word string) bool {
	// 定义一个深度优先搜索的递归函数 dfs
	var dfs func(x, y, l int) bool
	dfs = func(x, y, l int) bool {
		// 检查坐标 (x, y) 是否越界或当前字符是否不匹配
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] != word[l] {
			return false
		}
		// 如果已经匹配到 word 的最后一个字符，则返回 true
		if l == len(word)-1 {
			return true
		}
		// 临时保存当前字符，并标记当前位置已访问
		temp := board[x][y]
		board[x][y] = '/'
		// 向四个方向递归搜索
		res := dfs(x-1, y, l+1) || dfs(x, y+1, l+1) || dfs(x+1, y, l+1) || dfs(x, y-1, l+1)
		// 恢复当前位置的字符
		board[x][y] = temp
		return res
	}

	// 遍历二维数组的每一个元素，以每个元素为起点进行搜索
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 如果从 (i, j) 出发可以搜索到 word，则返回 true
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	// 如果没有找到匹配的字符串，返回 false
	return false
}
