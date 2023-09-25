package main

import "testing"

func BenchmarkC(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairs1(30)
		}
	})
	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairs2(30)
		}
	})
	b.Run("3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			climbStairs3(30)
		}
	})

}
