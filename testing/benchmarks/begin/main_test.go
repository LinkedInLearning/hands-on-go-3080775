// testing/benchmarks/begin/main_test.go
package main

import "testing"

// write a benchmark for sum
func BenchmarkSum(b *testing.B){
	for i := 0; i < b.N; i++ {
		sum(1, 2, 3)
	}
} 


// write a benchmark for sumAny
func BenchmarkSumAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAny([]interface{}{1,2,3})
	}
}
