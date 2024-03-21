package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}

// 快排 + 三数折中法取基准值
func sortArray(nums []int) []int {

	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left >= right {
			return
		}

		// 1、三数折中法取基准值（较其他随机选举法有性能优化）
		// 目标：最终，num[left] 为 num[left]、num[right]、num[mid] 三者中的中位数
		// 方法：前面两步先排序，排序好之后利用 num[left] 与 num[mid] 进行一次比较，就可以得出中位数
		mid := (left + right) / 2

		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[mid] > nums[right] {
			nums[mid], nums[right] = nums[right], nums[mid]
		}
		if nums[left] < nums[mid] {
			nums[left], nums[mid] = nums[mid], nums[left]
		}
		// stand 为中位数
		stand := nums[left]

		i, j := left, right
		for i < j {
			// 2、从右往左，找到第一个小于基准值的点
			for i < j && nums[j] >= stand {
				j--
			}
			// 3、从左往右，找到第一个大于基准值的点
			for i < j && nums[i] <= stand {
				i++
			}
			// 4、交换值
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 5、将基准值放到正确位置
		nums[left], nums[i] = nums[i], nums[left]
		// 6、重新定义基准节点，递归调用
		quickSort(left, i-1)
		quickSort(i+1, right)
	}

	quickSort(0, len(nums)-1)

	return nums
}
