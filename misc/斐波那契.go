package main

import "fmt"

func main() {
	fmt.Println(fibonacci(3))
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
