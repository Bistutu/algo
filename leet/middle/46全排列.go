package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	size := len(nums)

	// res 存储结果
	var res [][]int
	currentState := make([]int, 0, size)
	visited := make([]bool, size)
	var dfs func()

	dfs = func() {
		if len(currentState) == size {
			temp := make([]int, size)
			copy(temp, currentState)
			res = append(res, temp)
		}
		for i := 0; i < size; i++ {
			// 如果没有访问过
			if !visited[i] {
				visited[i] = true
				currentState = append(currentState, nums[i])
				dfs()
				visited[i] = false
				currentState = currentState[:len(currentState)-1]
			}
		}
	}

	dfs()
	return res
}
