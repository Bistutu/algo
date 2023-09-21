package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 2, 3, 4}))
}

func majorityElement(nums []int) int {
	length := len(nums)
	half := length / 2
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v > half {
			return k
		}
	}
	return 0
}

// 摩尔投票法（只要与我投的票不一样，就销毁）-> 我是拥有票最多的人，于是最后我肯定会胜出
func majorityElement2(nums []int) int {
	var target, count int
	for _, v := range nums {
		if count == 0 {
			target = v
		}
		if target != v {
			count--
		} else {
			count++
		}
	}
	return target
}
