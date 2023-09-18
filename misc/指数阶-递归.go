package main

import "fmt"

func main() {
	fmt.Println(powCountRecur(4))
}

func powCountRecur(n int) int {
	if n == 1 {
		return 1
	}
	return powCountRecur(n-1) + powCountRecur(n-1) - 1
}
