package main

import (
	"fmt"
)

type Stack []int

func main() {
	stack := make(Stack, 0)
	// push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	// peek
	fmt.Println(stack[len(stack)-1])
	// pop
	popValue := stack[len(stack)-1]
	fmt.Println(popValue)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
