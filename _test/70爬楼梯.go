package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	println(climbStairs1(30))
	fmt.Println(time.Now().Sub(now).String())
	now = time.Now()
	println(climbStairs2(30))
	fmt.Println(time.Now().Sub(now).String())
	now = time.Now()
	println(climbStairs3(30))
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

// 记忆化存储
func climbStairs3(n int) int {
	mem := make([]int, n+1)
	for k := range mem {
		mem[k] = -1 // -1 表示暂无记录
	}
	return handler(n, mem)
}

func handler(n int, mem []int) int {
	if n <= 2 {
		return n
	}
	if mem[n] != -1 {
		return mem[n]
	}
	count := handler(n-1, mem) + handler(n-2, mem)
	mem[n] = count
	return count
}
