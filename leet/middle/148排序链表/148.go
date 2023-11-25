package main

import (
	"fmt"
	"sort"
)

func sortList(head *ListNode) *ListNode {
	m := make([]*ListNode, 0)
	for head != nil {
		m = append(m, head)
		head = head.Next
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].Val < m[j].Val
	})
	temp := &ListNode{}
	cur := temp
	for i := 0; i < len(m); i++ {
		cur.Next = m[i]
		cur = m[i]
	}
	cur.Next = nil
	return temp.Next
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	sortedHead := sortList(head)
	for sortedHead != nil {
		fmt.Print(sortedHead.Val, " ")
		sortedHead = sortedHead.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
