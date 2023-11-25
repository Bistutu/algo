package main

import "fmt"

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	//n3 := &ListNode{Val: 3}
	//n4 := &ListNode{Val: 4}
	//n5 := &ListNode{Val: 5}
	n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	node := removeNthFromEnd(n1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
