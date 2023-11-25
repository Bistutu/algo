package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
	//people = [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}
	//fmt.Println(reconstructQueue(people))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, len(people))

	for _, v := range people {
		if len(res) <= v[1] {
			res = append(res, v)
		} else {
			copy(res[v[1]+1:], res[v[1]:])
			res[v[1]] = v
		}
	}
	return res
}
