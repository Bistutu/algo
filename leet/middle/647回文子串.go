package main

import "fmt"

func countSubstrings(s string) int {
	size := len(s)
	count := 0
	for center := 0; center < size*2-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < size && s[left] == s[right] {
			count++
			left--
			right++
		}
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("abcba")) // 输出应为 3
}
