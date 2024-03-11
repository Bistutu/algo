package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	rs := make([]string, 0)
	cur := make([]rune, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if len(cur) == 2*n {
			rs = append(rs, string(cur))
			return
		}
		if right > left { // 剪枝
			return
		}
		if left < n {
			cur = append(cur, '(')
			dfs(left+1, right)
			cur = cur[:len(cur)-1]
		}
		if right < left {
			cur = append(cur, ')')
			dfs(left, right+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, 0)

	return rs
}
