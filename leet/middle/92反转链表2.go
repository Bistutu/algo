package main

import (
	"fmt"
)

func main() {
	// TODO miankang.chen
	//[1,2,3,4,5], left = 2, right = 4
	//l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	l := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	node := reverseBetween(l, 1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	for i := left - 1; i < right; i++ {

	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
