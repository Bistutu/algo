package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

// 判断是否可以修完课程
// 秘诀是学会“拓扑排序”算法，知道如何操作
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1、查度、建立边的关系表
	inDegree := make(map[int]int, numCourses)
	graph := make(map[int][]int, numCourses)
	for _, v := range prerequisites { // v 里面就两个元素
		inDegree[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	// 初始化队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 检测拓扑关系
	count := 0
	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]
		count++
		for _, v := range graph[temp] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return count == numCourses
}
