package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}

func minMeetingRooms(intervals [][]int) int {
	count := 0
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	for k, v := range intervals {
		start[k] = v[0]
		end[k] = v[1]
	}

	sort.Ints(start)
	sort.Ints(end)

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
