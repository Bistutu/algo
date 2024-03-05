package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("aaaaaaaaaa", "aaaaaaaaaaaaa"))
	//fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	rs := make([]int, 0)
	n1, n2 := len(s), len(p)
	if n2 > n1 {
		return rs
	}
	var sm, pm [26]int
	for i := 0; i < n2; i++ {
		sm[s[i]-'a']++
		pm[p[i]-'a']++
	}
	if checkSame(sm, pm) {
		rs = append(rs, 0)
	}

	for i := n2; i < n1; i++ {
		sm[s[i]-'a']++
		sm[s[i-n2]-'a']--
		if checkSame(sm, pm) {
			rs = append(rs, i-n2+1)
		}
	}
	return rs
}

func checkSame(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
