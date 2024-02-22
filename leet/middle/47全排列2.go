package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}

func permuteUnique(nums []int) [][]int {
	size := len(nums)

	res := make([][]int, 0)
	path := make([]int, 0, size)
	visited := make([]bool, size)

	var dfs func()
	dfs = func() {
		if len(path) == size {
			t := make([]int, size)
			copy(t, path)
			res = append(res, t)
		}
		duplicate := make(map[int]struct{})
		for i := 0; i < size; i++ {
			// 如果没有访问过
			if _, ok := duplicate[nums[i]]; !ok && !visited[i] {
				visited[i] = true
				duplicate[nums[i]] = struct{}{}
				path = append(path, nums[i])
				dfs()
				// 回退
				visited[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return res
}
