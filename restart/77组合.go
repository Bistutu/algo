package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	rs := make([][]int, 0)
	cur := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {
		if len(cur) == k {
			t := make([]int, len(cur))
			copy(t, cur)
			rs = append(rs, t)
			return // 记得 return
		}
		for i := start; i <= n; i++ {
			cur = append(cur, i)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(1)
	return rs
}
