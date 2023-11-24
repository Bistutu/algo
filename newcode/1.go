package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(1000000000))
}

func findNthDigit(n int) int {
	length := 1 // 当前数字的长度
	count := 9  // 在当前长度下，有多少个数字
	start := 1  // 当前长度的第一个数字

	// 确定n落在哪个长度区间
	for n > length*count {
		n -= length * count
		length++
		count *= 10
		start *= 10
	}

	start += (n - 1) / length // 确定n落在的那个数字
	s := strconv.Itoa(start)
	return int(s[(n-1)%length] - '0') // 计算并返回n的具体数值
}
