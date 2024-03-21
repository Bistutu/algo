package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	cur := head
	// 注：dummy 无论何时何地都是“头节点”，每轮循环都一样
	for cur != nil {
		next := cur.Next // 保存现场
		prev := dummy
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		cur.Next = prev.Next
		prev.Next = cur

		cur = next
	}
	return dummy.Next
}

func main() {
	// 创建一个单向链表
	head := &ListNode{Val: 4, Next: nil}
	node1 := &ListNode{Val: 2, Next: nil}
	node2 := &ListNode{Val: 1, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	// 对链表进行插入排序
	sortedHead := insertionSortList(head)

	// 打印排序后的链表
	fmt.Println("排序后的链表：")
	printList(sortedHead)
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
