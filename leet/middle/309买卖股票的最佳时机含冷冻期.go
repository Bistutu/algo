package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}), 3)
}

func maxProfit(prices []int) int {
	days := len(prices)
	if days == 0 {
		return 0
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
