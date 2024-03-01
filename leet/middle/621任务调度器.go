package main

import (
	"fmt"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}

// 理解，但是如果重头写则无招
func leastInterval(tasks []byte, n int) int {
	// 1、初始化任务
	freq := make(map[byte]int, len(tasks))
	maxFreq, maxCount := 0, 0 // 最高频率任务的频率、以及相同最高频率的任务数量
	for _, v := range tasks {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
			maxCount = 1
		} else if freq[v] == maxFreq {
			maxCount++
		}
	}

	// 2、开始运作（计算所需最少时间的核心逻辑）
	partCount := maxFreq - 1                        // 最高频率任务直接的间隔数量
	partLength := n - (maxCount - 1)                // 每个间隔的长度
	emptySlo := partCount * partLength              // 空闲插槽的总数
	availableTasks := len(tasks) - maxFreq*maxCount // 还未安排的任务数量
	idles := max(0, emptySlo-availableTasks)        // 计算要插入的空闲时间
	return len(tasks) + idles
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
