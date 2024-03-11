package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	rs := make([][]string, 0)
	m := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] > bytes[j]
		})
		str := string(bytes)
		m[str] = append(m[str], v)
	}
	for _, v := range m {
		rs = append(rs, v)
	}
	return rs
}
