package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}

func minMeetingRooms(intervals [][]int) int {
	size := len(intervals)
	if size == 0 {
		return 0
	}
	// 开始时间列表和结束时间列表
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	for i := 0; i < size; i++ {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}

	sort.Ints(start)
	sort.Ints(end)

	count := 0
	index := 0
	for _, v := range start {
		if v >= end[index] {
			index++
		} else {
			count++
		}
	}
	return count
}
