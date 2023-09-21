package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(time.Now().Sub(now).String())
	now = time.Now()
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(time.Now().Sub(now).String())
}

// 暴力解法
func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	ints := make([]int, size)
	k := 0
	for i := 0; i < size; i++ {
		i2 := i + 1
		if i >= 1 && temperatures[i] > temperatures[i-1] {
			i2 = k
		}
		for j := i2; j < size; j++ {
			if temperatures[i] < temperatures[j] {
				ints[i] = j - i
				k = j
				break
			}
		}
	}
	return ints
}

func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0) // 用切片模拟栈

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			res[lastIndex] = i - lastIndex
			stack = stack[:len(stack)-1] // pop
		}
		stack = append(stack, i) // push
	}

	return res
}
