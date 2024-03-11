package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {
	sm := make(map[rune]int, len(s))
	tm := make(map[rune]int, len(t))
	if len(s) != len(t) {
		return false
	}
	for _, v := range s {
		sm[v]++
	}
	for _, v := range t {
		tm[v]++
	}
	for k := range sm {
		if sm[k] != tm[k] {
			return false
		}
	}
	return true
}
