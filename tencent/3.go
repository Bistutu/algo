package main

import "fmt"

func countTriplets(a, b, c []int, target int) int {
	count := 0
	m := make(map[int]int, len(a))
	for _, v1 := range a {
		for _, v2 := range b {
			m[v1+v2]++
		}
	}
	for _, v := range c {
		if val, ok := m[target-v]; ok {
			count += val
		}
	}
	return count
}

func main() {
	a := []int{20, 30, 40}
	b := []int{30, 20, 10}
	c := []int{6, 16, 26}
	target := 66

	fmt.Println(countTriplets(a, b, c, target))
}
