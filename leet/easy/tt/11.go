package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// 创建环：将第五个节点的 Next 指向第三个节点
	head.Next.Next.Next.Next.Next = head.Next.Next
	fmt.Println(hasCycle(head))
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast != nil && slow.Val == fast.Val { // 比较指针而不是值
			return true
		}
	}
	return false
}
