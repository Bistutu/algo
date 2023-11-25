package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		temp := []byte(v)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] > temp[j]
		})
		m[string(temp)] = append(m[string(temp)], v)
	}
	rs := make([][]string, 0, len(m))
	for _, v := range m {
		rs = append(rs, v)
	}
	return rs
}
