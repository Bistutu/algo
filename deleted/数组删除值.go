package main

import "fmt"

func main() {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i
	}
	// 删除 1 + 挪动
	for i := 1; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	array[len(array)-1] = 0
	for k, v := range array {
		fmt.Println(k, v)
	}
}
