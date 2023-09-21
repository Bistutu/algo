package main

func main() {
	getIntersectionNode(&ListNode{Val: 1}, &ListNode{Val: 2})
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
