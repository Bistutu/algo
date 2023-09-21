package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		for _, v := range binary {
			if v == '1' {
				result[i]++
			}
		}
	}
	return result
}
