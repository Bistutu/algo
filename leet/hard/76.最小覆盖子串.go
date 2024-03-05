package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

// 秘诀：初始左右边界均为 0，不断移动右边界，当遇到符合条件时尝试移动左边界，并最终取最小符合组
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	// 统计 t 中各字符出现的次数
	tm := make(map[rune]int)
	for _, v := range t {
		tm[v]++
	}

	// 定义窗口内字符频率 map，window map
	wm := make(map[rune]int)

	// 滑动窗口算法左右边界，在此场景下初始值均为 0
	left, right := 0, 0
	// 记录当前窗口已包含的 t 中的字符数量，以及需要的字符总数
	have, need := 0, len(tm)
	// 记录最小覆盖子串的其实索引及长度
	start, length := 0, math.MaxInt

	for right < len(s) {
		rc := rune(s[right]) // right char
		wm[rc]++

		// 如果当前字符是 t 中的字符，且数量刚好达到需求，则更新 have
		if wm[rc] == tm[rc] {
			have++
		}
		right++

		// 当窗口中已经包含 t 中所有字符的时候，尝试收缩左边界
		for have == need {
			// 尝试获取最小值
			if right-left < length {
				start = left
				length = right - left
			}
			// 尝试向左移动字符（并更新相关计数）
			lc := rune(s[left])
			wm[lc]--
			// 如果移动的是必要字符
			if tm[lc] > 0 && wm[lc] < tm[lc] {
				have--
			}
			left++
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
