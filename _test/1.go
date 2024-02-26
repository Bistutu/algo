package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countPrimes(10)) // 输出小于10的质数数量
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
