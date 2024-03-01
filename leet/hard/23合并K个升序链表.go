package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// 创建链表 [[1,4,5],[1,3,4],[2,6]]
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}
	lists := []*ListNode{list1, list2, list3}

	// 合并链表
	merged := mergeKLists(lists)

	// 打印合并后的链表
	for merged != nil {
		fmt.Printf("%d ", merged.Val)
		merged = merged.Next
	}
}

// 小跟堆
func mergeKLists(lists []*ListNode) *ListNode {

	m := &minHeap{}
	heap.Init(m)

	// 添加至堆
	for _, v := range lists {
		for v != nil {
			heap.Push(m, v)
			v = v.Next
		}
	}

	node := &ListNode{}
	temp := node

	// 出堆构建链表
	for m.Len() != 0 {
		val := heap.Pop(m)
		node.Next = val.(*ListNode)
		node = val.(*ListNode)
	}
	node.Next = nil

	return temp.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(*ListNode))
}
func (h *minHeap) Pop() any {
	temp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return temp
}

// 暴力解法
func mergeKLists2(lists []*ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	for _, v := range lists {
		for v != nil {
			nodes = append(nodes, v)
			v = v.Next
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	root := &ListNode{}
	temp := root

	for _, v := range nodes {
		root.Next = v
		root = v
	}
	root.Next = nil

	return temp.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
