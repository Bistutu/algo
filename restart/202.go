package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	// 为了避免 loop，添加 visited 记录
	visited := make(map[int]bool)

	str := strconv.Itoa(n)
	sum := 0
	for !visited[sum] {
		visited[sum] = true
		sum = 0
		for _, v := range str {
			num, _ := strconv.Atoi(string(v))
			sum += num * num
		}
		if sum == 1 {
			return true
		}
		str = strconv.Itoa(sum)
	}
	return false
}
