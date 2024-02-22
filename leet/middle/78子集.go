package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	r := make([][]int, 0) // 初始化结果集

	// 使用二进制表示所有可能的组合
	for i := 0; i < 1<<len(nums); i++ { // 从0到2^len(nums)-1，遍历所有组合
		temp := make([]int, 0) // 初始化当前子集

		// 遍历nums中的每个元素
		for k, v := range nums {
			// i的第k位如果为1，表示nums中的第k个元素应该被包含在当前子集中
			if i>>k&1 > 0 {
				temp = append(temp, v) // 将元素添加到当前子集
			}
		}

		// 将当前子集添加到结果集中
		r = append(r, temp)
	}
	return r // 返回所有子集的集合
}
