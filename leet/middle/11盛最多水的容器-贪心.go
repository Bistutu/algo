package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		now := (j - i) * min(height[i], height[j])
		if now > max {
			max = now
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
