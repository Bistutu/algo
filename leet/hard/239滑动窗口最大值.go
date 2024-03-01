package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{} // 返回空切片，因为没有有效的窗口或输入数组
	}
	rs := make([]int, 0)
	h := &IntHeap{}
	heap.Init(h)
	// 先把 k-1 个元素加进去（代表再加入一个元素就应该能得到一个结果 rs）
	for i := 0; i < k-1; i++ {
		heap.Push(h, [2]int{nums[i], i})
	}
	for i := k - 1; i < len(nums); i++ {
		heap.Push(h, [2]int{nums[i], i})
		// 每次都循环清除堆中不属于当前窗口的元素！
		for (*h)[0][1] <= i-k {
			heap.Pop(h)
		}
		rs = append(rs, (*h)[0][0])
	}
	return rs
}

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
