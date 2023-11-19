package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cuttingBamboo(5)) // 输出 81
}

func cuttingBamboo(n int) int {
	if n <= 3 {
		return 1 * (n - 1)
	}
	count := n / 3
	remainder := n % 3
	if remainder == 1 {
		count--
		return int(math.Pow(3, float64(count))) * 4
	}
	if remainder == 2 {
		return int(math.Pow(3, float64(count))) * 2
	}
	return int(math.Pow(3, float64(count)))
}
