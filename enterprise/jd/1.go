package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count) // 场次

	nums := make([][]rune, 3) // 每局
	for i := 0; i < count; i++ {

		a, b := false, false

		for j := 0; j < 3; j++ {
			nums[j] = make([]rune, 3)
			fmt.Scanf("%c%c%c\n", &nums[j][0], &nums[j][1], &nums[j][2])
		}
		// 判断是谁获胜
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// 跳过空节点
				if nums[i][j] == '.' {
					continue
				}

				if ok := isSurround(nums, i, j); ok {
					fmt.Println(string(nums[i][j]))
					switch nums[i][j] {
					case 'o':
						b = true
					case '*':
						a = true
					}
				}
			}
		}
		if a && b {
			fmt.Println("draw")
		} else if a {
			fmt.Println("kou")
		} else {
			fmt.Println("yukari")
		}
	}

}

func isSurround(nums [][]rune, x, y int) bool {
	if (x == 1) && nums[x-1][y] == nums[x+1][y] && nums[x-1][y] != nums[x][y] && nums[x-1][y] != '.' {
		return true
	}

	if (y == 1) && nums[x][y-1] == nums[x][y+1] && nums[x][y-1] != nums[x][y] && nums[x-1][y] != '.' {
		return true
	}
	return false
}

func echo() {
	fmt.Println("asa")
}
