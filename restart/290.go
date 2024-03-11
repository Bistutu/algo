package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	n1, n2 := len(pattern), len(strs)
	if n1 != n2 {
		return false
	}
	if n1 == 1 {
		return true
	}

	pm := make(map[byte]string)
	sm := make(map[string]byte)

	for k := range pattern {
		pp, ss := pattern[k], strs[k]
		if len(pm[pp]) > 0 && pm[pp] != ss || sm[ss] > 0 && sm[ss] != pp {
			return false
		}
		pm[pp] = ss
		sm[ss] = pp
	}

	return true
}
