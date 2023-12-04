package main

import "fmt"

func main() {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))

}

func canAttendMeetings(intervals [][]int) bool {
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] >= intervals[i+1][0] && intervals[i][0] < intervals[i+1][1] {
			return false
		}
		if intervals[i+1][0] > intervals[i][0] && intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}

	return true
}
