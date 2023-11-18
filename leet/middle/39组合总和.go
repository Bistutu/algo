package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var current []int
	var dfs func(start int)
	size := len(candidates)

	dfs = func(start int) {
		switch {
		case sum(current) == target:
			temp := make([]int, len(current))
			copy(temp, current)
			res = append(res, temp)
		case sum(current) > target: // 剪枝
			return
		}
		for i := start; i < size; i++ {
			current = append(current, candidates[i])
			dfs(i)
			current = current[:len(current)-1]
		}
	}
	dfs(0)
	return res
}

func sum(nums []int) int {
	count := 0
	for _, v := range nums {
		count += v
	}
	return count
}
