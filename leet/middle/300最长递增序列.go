package main

import "fmt"

func main() {
	var n interface{} = &P{}
	v, ok := n.(int)
	fmt.Println(v, ok)
}

type P struct {
	msg string
}

func lengthOfLIS(nums []int) int {
	return 0
}
