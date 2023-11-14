package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}))
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	ints := make([]float64, 0, 6)
	ints = append(ints, distance(p1, p2))
	ints = append(ints, distance(p2, p3))
	ints = append(ints, distance(p3, p4))
	ints = append(ints, distance(p4, p1))
	ints = append(ints, distance(p1, p3))
	ints = append(ints, distance(p2, p4))
	sort.Float64s(ints)
	if ints[0] != 0 && ints[4] != 0 && ints[0] == ints[1] && ints[1] == ints[2] && ints[2] == ints[3] && ints[4] == ints[5] {
		return true
	}
	return false
}

func distance(a, b []int) float64 {
	return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
}
