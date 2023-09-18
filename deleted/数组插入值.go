package main

import "fmt"

func main() {
	var array [5]int
	for i := 0; i < 4; i++ {
		array[i] = i
	}
	// 我现在想要在 0 后面插入一个数 5
	for i := len(array) - 1; i > 1; i-- {
		array[i] = array[i-1]
	}
	array[1] = 5
	for k, v := range array {
		fmt.Println(k, v)
	}
}
