package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(backtrackSubsetSumINaive([]int{3, 4, 5}, 9))
}

func backtrackSubsetSumINaive(nums []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	duplicate := make(map[string]struct{}) // 去除重复元素

	var dfs func(cur int)
	dfs = func(cur int) {
		if cur > target {
			return
		}

		if cur == target {
			sort.Ints(path)
			s := fmt.Sprint(path)
			if _, ok := duplicate[s]; !ok {
				duplicate[s] = struct{}{}
				t := make([]int, len(path))
				copy(t, path)
				res = append(res, t)
			}
		}
		for _, v := range nums {
			path = append(path, v)
			dfs(cur + v)
			// 回退
			path = path[:len(path)-1]
		}
	}
	dfs(0)

	return res
}
