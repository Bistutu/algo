package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	numbers := addTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	)
	for numbers != nil {
		fmt.Println(numbers.Val)
		numbers = numbers.Next
	}
	// l1=[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// l2=[5,6,4]
	/*l1 := &ListNode{1, nil}
	current := l1
	for i := 0; i < 29; i++ {
		current.Next = &ListNode{0, nil}
		current = current.Next
	}
	current.Next = &ListNode{1, nil}

	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	numbers2 := addTwoNumbers(l1, l2)

	for numbers2 != nil {
		fmt.Printf("%d ", numbers2.Val)
		numbers2 = numbers2.Next
	}*/
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 暴力解法
	num1 := make([]int, 0)
	num2 := make([]int, 0)
	count1 := 0
	count2 := 0
	for l1 != nil {
		num1 = append(num1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		num2 = append(num2, l2.Val)
		l2 = l2.Next
	}
	for i := len(num1) - 1; i >= 0; i-- {
		count1 += num1[i] * int(math.Pow10(i))
	}
	for i := len(num2) - 1; i >= 0; i-- {
		count2 += num2[i] * int(math.Pow10(i))
	}
	count := count1 + count2
	itoa := strconv.Itoa(count)
	l3 := &ListNode{}
	l4 := l3
	for i := len(itoa) - 1; i >= 0; i-- {
		u := itoa[i] - 48
		l3.Val = int(u)
		if i != 0 {
			l3.Next = &ListNode{}
			l3 = l3.Next
		}
	}
	return l4
}
