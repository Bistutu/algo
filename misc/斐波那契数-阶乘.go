package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return factorial(n-1) + factorial(n-2)
	}
}
