package main

import "testing"

func BenchmarkC(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			climbStairs1(30)
		}
	})
}
