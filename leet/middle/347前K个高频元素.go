package main

import (
	"fmt"
	"sort"
)

type T struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	t := make([]*T, 0, len(m))
	for k, v := range m {
		t = append(t, &T{
			num:   k,
			count: v,
		})
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].count > t[j].count
	})
	rs := make([]int, 0, k)
	for i := 0; i < k; i++ {
		rs = append(rs, t[i].num)
	}

	return rs
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2}, 2))
}
