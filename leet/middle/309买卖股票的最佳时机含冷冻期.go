package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}), 3)
}

// 前提：只能同时买卖一支股票，不能同时持有多支股票
func maxProfit(prices []int) int {
	day := len(prices)
	if day <= 1 {
		return 0
	}
	// 初始化三个数组
	hold := make([]int, day)     // 持有股票状态下的最大利润，表示在当天结束时你持有股票
	frozen := make([]int, day)   // 冷冻期状态下的最大利润，表示当天结束时，你卖出了股票并处于冷冻期
	unfrozen := make([]int, day) // 非冷冻状态下的最大利润，表示当天结束时，你没有股票且不处于冷冻期

	// PS. 所以其实最后的最大利润肯定不是 hold[day]，因为最后一天我们一定要卖出股票！

	hold[0] = -prices[0]
	frozen[0] = 0
	unfrozen[0] = 0

	// 从第二天开始，遍历每一天的价格（同时更新三种状态下的每一天最大利润）
	for i := 1; i < day; i++ {
		// 1、继续持有 或 从非冷冻期购买
		hold[i] = max(hold[i-1], unfrozen[i-1]-prices[i])
		// 2、今天卖出股票
		frozen[i] = hold[i-1] + prices[i]
		// 继续非冷冻期 或 进入冷冻期
		unfrozen[i] = max(unfrozen[i-1], frozen[i-1])
	}
	// 返回最后一天结束时，冷冻期和非冷冻期下的最大利润者
	return max(frozen[day-1], unfrozen[day-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
