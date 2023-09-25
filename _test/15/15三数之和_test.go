package main

import (
	"testing"
)

func BenchmarkOne(b *testing.B) {
	nums := []int{-2, 0, 1, 1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		threeSum(nums)
	}
}
