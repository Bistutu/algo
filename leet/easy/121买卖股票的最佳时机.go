package main

import "fmt"

func main() {
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit3(prices []int) int {
	min := prices[0]
	maxProfile := 0
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > maxProfile {
			maxProfile = v - min
		}
	}
	return maxProfile
}

// 暴力解法
func maxProfit(prices []int) int {
	length := len(prices)
	pros := make([]int, length)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if pros[i] < prices[j]-prices[i] {
				pros[i] = prices[j] - prices[i]
			}
		}
	}
	maxProfit := 0
	for _, v := range pros {
		if v > maxProfit {
			maxProfit = v
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var minPrice = prices[0]
	maxProfit := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}
	return maxProfit
}
