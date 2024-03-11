package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	rs := make([][]int, 0)
	visited := make(map[int]bool, len(nums))
	cur := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(cur) == len(nums) {
			t := make([]int, len(cur))
			copy(t, cur)
			rs = append(rs, t)
			return
		}
		for _, v := range nums {
			if visited[v] {
				continue
			}
			visited[v] = true
			cur = append(cur, v)
			dfs()
			visited[v] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs()
	return rs
}
