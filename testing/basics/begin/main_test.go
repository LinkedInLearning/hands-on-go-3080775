// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for sum
func TestSum(t *testing.T) {
	result := sum(1, 2, 3)
	want := 6
	if result != want {
		t.Errorf("result was %v and wanted is %v", result, want)
	}
}

// write a TestMain for setup and teardown
func TestMain(m *testing.M) {
	fmt.Println("***SETUP***")
	m.Run()
	fmt.Println("***TEARDOWN***")
}
