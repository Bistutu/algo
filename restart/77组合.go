package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	// 这里现在虽然不知道有多少个，但是 []int 肯定为 k
	rs := make([][]int, 0)

	return rs
}
