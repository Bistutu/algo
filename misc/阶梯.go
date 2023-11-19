package main

import "fmt"

func main() {
	fmt.Println(climbingStairsDFS(5))
}
func climbingStairsDFS(n int) int {
	memo := make(map[int]int)

	var dfs func(n int) int
	dfs = func(n int) int {
		if n <= 2 {
			return n
		}
		if v, ok := memo[n]; ok {
			return v
		}
		memo[n] = dfs(n-1) + dfs(n-2)
		return memo[n]
	}
	return dfs(n)
}
