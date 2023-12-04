package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(nums))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rs := make([][]int, 0)
	for _, v := range intervals {
		if len(rs) == 0 || rs[len(rs)-1][1] < v[0] {
			rs = append(rs, v)
		} else {
			rs[len(rs)-1][1] = max(rs[len(rs)-1][1], v[1])
		}
	}
	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
