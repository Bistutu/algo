package main

import "fmt"

func main() {
	// 指数阶
	fmt.Println(powCount(3))
}

func powCount(n int) int {
	count, base := 0, 1
	for i := 0; i < n; i++ {
		for i := 0; i < base; i++ {
			count++
		}
		base *= 2
	}
	return count
}
