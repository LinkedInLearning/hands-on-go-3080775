package main

import "golang.org/x/exp/constraints"

type numeric interface {
	constraints.Integer | constraints.Float
}

func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func sumAny(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}
