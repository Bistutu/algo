package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	println(climbStairs2(3))
	fmt.Println(time.Now().Sub(now).String())
}

// 递归（时间：2^n，效率极低）
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}

// 动态规划
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	pre, cur := 1, 2
	for i := 3; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
