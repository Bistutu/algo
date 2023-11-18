package main

import "fmt"

func main() {
	// [3,2,0,-4]
	fmt.Println(hasCycle2(&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}))
	// [1,2]
	fmt.Println(hasCycle(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
	// [1]
	fmt.Println(hasCycle(&ListNode{Val: 1}))
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		} else {

			m[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	slow := head.Next
	fast := head.Next.Next
	for slow != nil && fast != nil {
		if slow.Val == fast.Val {
			return true
		}
		slow = head.Next
		fast = head.Next.Next
	}
	return false
}
