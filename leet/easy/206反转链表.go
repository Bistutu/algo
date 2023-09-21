package main

func main() {
	// 链表 = 1,2,3,4,5
	l := reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

// 暴力解法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ints := make([]int, 0)
	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}
	first := &ListNode{Val: ints[len(ints)-1]}
	temp := first

	for i := len(ints) - 2; i >= 0; i-- {
		temp.Next = &ListNode{Val: ints[i]}
		temp = temp.Next
	}
	return first
}
