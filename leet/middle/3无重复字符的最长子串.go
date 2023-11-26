package main

import "fmt"

func main() {
	examples := []string{"abba", "bbbbb", "pwwkew"}
	for _, example := range examples {
		fmt.Println(lengthOfLongestSubstring(example))
	}
}

func lengthOfLongestSubstring(s string) int {
	site := make(map[byte]int)
	maxLength := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		if v, ok := site[s[right]]; ok && v > left {
			left = v + 1
		}
		site[s[right]] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}
