package main

import "fmt"

func main() {
	fmt.Println(factorial3(5, 1))
}

func factorial3(n, m int) int {
	if n == 0 {
		return m
	}
	return factorial3(n-1, n*m)
}
