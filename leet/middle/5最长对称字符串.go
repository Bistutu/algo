package main

import "fmt"

func longestPalindrome(s string) string {
	// 如果字符串长度小于1，则直接返回空字符串
	if len(s) < 1 {
		return ""
	}

	var handle func(left int, right int) int
	handle = func(left int, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}
	var max func(x int, y int) int
	max = func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 初始化回文子串的起始和结束索引，遍历字符串，以每个字符为中心扩展寻找回文子串
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 以字符 s[i]/s[i] 和 s[i+1]之间 为中心的最长回文子串长度
		len1 := handle(i, i)
		len2 := handle(i, i+1)

		// 选择两者中较长的回文长度
		len := max(len1, len2)

		// 如果找到了更长的回文子串，则更新起始和结束索引
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func main() {
	fmt.Println("最长对称字符串为:", longestPalindrome("123abccba123"))
}
