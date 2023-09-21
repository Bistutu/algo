package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
