package main

import "fmt"

func main() {
	fmt.Println(factorial2(4))
}

func factorial2(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial2(n-1)
}
