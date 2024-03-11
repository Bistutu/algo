package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	rs := make([][]int, 0)
	cur := make([]int, 0)

	// 带上 start 就不会走回头路了
	var dfs func(start int)
	dfs = func(start int) {
		now := sum(cur)
		if now == target {
			t := make([]int, len(cur))
			copy(t, cur)
			rs = append(rs, t)
			return
		} else if now > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			dfs(i)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return rs
}

func sum(nums []int) int {
	count := 0
	for _, v := range nums {
		count += v
	}
	return count
}
