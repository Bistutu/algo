package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
	}

	for _, test := range tests {
		got := exist(test.board, test.word)
		if got != test.want {
			fmt.Printf("Failed for board: %v, word: %s, got: %t, want: %t\n", test.board, test.word, got, test.want)
		}
	}
}

func exist(board [][]byte, word string) bool {
	n, m, target := len(board), len(board[0]), len(word)
	visited := make([][]bool, n)
	for k := range visited {
		visited[k] = make([]bool, m)
	}

	var dfs func(i, j, cur int) bool
	dfs = func(i, j, cur int) bool {
		if cur == target {
			return true
		}

		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != word[cur] || visited[i][j] {
			return false
		}
		visited[i][j] = true
		rs := dfs(i+1, j, cur+1) || dfs(i-1, j, cur+1) || dfs(i, j+1, cur+1) || dfs(i, j-1, cur+1)
		visited[i][j] = false
		return rs
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
