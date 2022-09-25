package main

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2, 3)
	}
}

func BenchmarkSumAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAny([]interface{}{1, 2, 3})
	}
}
