// testing/table-driven/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, -1},
		{"three", []int{1, 2, 3}, 7},
	}

	// range over the tests and run them as subtests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if sum(test.input...) != test.want {
				t.Error("Test failed")
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("****")
	m.Run()
	fmt.Println("***")
}
