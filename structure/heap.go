package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := &intHeap{}
	heap.Init(h)
	heap.Push(h, 1)
	heap.Push(h, 3)
	heap.Push(h, 0)
	heap.Push(h, -2)
	fmt.Println(h.Peek())
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
}

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *intHeap) Pop() any {
	temp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return temp
}

func (h *intHeap) Peek() any {
	return (*h)[0]
}
