package main

import (
	"fmt"
)

func main() {
	space := []int{2, 2, 3, 1, 5, 2}
	value := []int{2, 3, 1, 5, 4, 3}
	fmt.Println(knapsackDP(space, value, 1))
}

func knapsackDP(wgt, val []int, cap int) int {
	n := len(wgt)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, cap+1)
	}
	return 0
}
