package main

import (
	"container/heap"
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
	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2}, 2))
}

func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	h := &myHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, T{
			num:   k,
			count: v,
		})
	}
	ints := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ints = append(ints, heap.Pop(h).(T).num)
	}
	return ints
}

type myHeap []T

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	return m[i].count > m[j].count
}

func (m myHeap) Swap(i, j int) {
	t := m[i]
	m[i] = m[j]
	m[j] = t
}

func (m *myHeap) Push(val any) {
	*m = append(*m, val.(T))
}

func (m *myHeap) Pop() any {
	t := (*m)[m.Len()-1]
	*m = (*m)[0 : (*m).Len()-1]
	return t
}
