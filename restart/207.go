package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
}

// 精髓就是“会定义初始状态”，即根据“拓扑排序”来
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建拓扑排序所需变量
	in := make(map[int]int, numCourses)      // 入度统计
	graph := make(map[int][]int, numCourses) // 边
	for _, v := range prerequisites {
		in[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	// 构建初始队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ { // todo 注意这里判断的为 i < numCourses
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 开始拓扑排序的操作，统计已经拓扑到的数字数量
	count := 0
	for len(queue) != 0 {
		count++
		// 弹出数据、删除边
		t := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, v := range graph[t] {
			in[v]--
			if in[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return count == numCourses
}
