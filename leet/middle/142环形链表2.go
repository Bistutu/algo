package main

import "fmt"

func main() {
	//head = [3,2,0,-4], pos = 1
	node := detectCycle(createLinkedListWithCycle([]int{3, 2, 0, -4}, 1))
	fmt.Println(node.Val)
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func createLinkedListWithCycle(values []int, pos int) *ListNode {
	var head, tail, cycleStart *ListNode

	for i, val := range values {
		newNode := &ListNode{Val: val}

		if i == 0 {
			head = newNode
		} else {
			tail.Next = newNode
		}

		if i == pos {
			cycleStart = newNode
		}

		tail = newNode
	}

	if cycleStart != nil {
		tail.Next = cycleStart
	}

	return head
}
