package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {
	runes := make(map[rune]int, len(magazine))
	for _, v := range magazine {
		runes[v]++
	}
	for _, v := range ransomNote {
		if time, ok := runes[v]; ok && time > 0 {
			runes[v]--
		} else {
			return false
		}
	}
	return true
}
