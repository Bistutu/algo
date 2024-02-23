package main

import (
	"fmt"
)

func main() {
	// 1、暴力搜索
	wgt := []int{10, 20, 30, 40, 50}
	val := []int{50, 120, 150, 210, 240}
	count, cap := 5, 50
	fmt.Println(knapsackDFS(wgt, val, count, cap))
	// 2、暴力搜索+记忆化
	mem := make([][]int, count)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, cap)
	}
	fmt.Println(knapsackDFSMem(wgt, val, mem, count, cap))
	// 3、动态规划
	fmt.Println(knapsackDP(wgt, val, count, cap))
	//fmt.Println(knapsackDP([]int{1, 2, 3}, []int{5, 11, 15}, 3, 4))

	// 特例：完全背包（可重复选取），其与 01 背包最大的区别就是状态转移方程发生了一丝改变，见：https://www.hello-algo.com/chapter_dynamic_programming/unbounded_knapsack_problem/#1
	fmt.Println(unboundedKnapsackDP(wgt, val, count, cap))
}

func unboundedKnapsackDP(wgt, val []int, count, cap int) int {
	dp := make([][]int, count+1)
	for i := 0; i < count+1; i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 1; i <= count; i++ {
		for j := 1; j <= cap; j++ {
			// 重量超出，没办法
			if wgt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 要么不选，要么就“刚刚好塞下”时的最大值
				// 对比 01 背包：状态转移中有一处从 i-1 变为 i，其余完全一致
				dp[i][j] = max(dp[i-1][j], dp[i][j-wgt[i-1]]+val[i-1])
			}
		}
	}
	return dp[count][cap]
}

// 动态规划
func knapsackDP(wgt, val []int, count, cap int) int {
	dp := make([][]int, count+1)
	for i := 0; i < count+1; i++ {
		dp[i] = make([]int, cap+1)
	}
	for i := 1; i <= count; i++ {
		for j := 1; j <= cap; j++ {
			// 重量超出，没办法
			if wgt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 要么不选，要么就“刚刚好塞下”时的最大值
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wgt[i-1]]+val[i-1])
			}
		}
	}
	return dp[count][cap]
}

// 暴力解法+记忆化
func knapsackDFSMem(wgt, val []int, mem [][]int, count, cap int) int {
	if count == 0 || cap == 0 {
		return 0
	}
	// 超过背包容量，只能不放入
	if wgt[count-1] > cap {
		return knapsackDFSMem(wgt, val, mem, count-1, cap)
	}
	if mem[count-1][cap-1] != 0 {
		return mem[count][cap]
	}
	// 计算放入与不放入，不放入时本次价值为 0，故不需要添加 val
	no := knapsackDFSMem(wgt, val, mem, count-1, cap)
	yes := knapsackDFSMem(wgt, val, mem, count-1, cap-wgt[count-1]) + val[count-1]

	rsp := max(no, yes)
	mem[count-1][cap-1] = rsp
	return rsp
}

// 暴力解法
func knapsackDFS(wgt, val []int, count, cap int) int {
	if count == 0 || cap == 0 {
		return 0
	}
	// 超过背包容量，只能不放入
	if wgt[count-1] > cap {
		return knapsackDFS(wgt, val, count-1, cap)
	}
	// 计算放入与不放入，不放入时本次价值为 0，故不需要添加 val
	no := knapsackDFS(wgt, val, count-1, cap)
	yes := knapsackDFS(wgt, val, count-1, cap-wgt[count-1]) + val[count-1]

	return max(no, yes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
