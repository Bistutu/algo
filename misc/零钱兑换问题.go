// 背包问题的变种：零钱兑换（完全背包）
package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	// 零钱兑换1
	fmt.Println(coinChange(coins, 11))
	// 零钱兑换2
	fmt.Println(change(5, coins))
}

// 求能凑齐硬币数的方案数量
func change(amount int, coins []int) int {
	size := len(coins)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= size; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < size+1; i++ {
		for j := 1; j < amount+1; j++ { // 空间
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[size][amount]
}

// 求能够凑齐的最少硬币数量，如无法凑齐则返回 -1
func coinChange(coins []int, amount int) int {
	size := len(coins)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 1; i < amount+1; i++ {
		dp[0][i] = amount + 1
	}

	for i := 1; i < size+1; i++ {
		for j := 1; j < amount+1; j++ { // 空间
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[size][amount] != amount+1 {
		return dp[size][amount]
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
