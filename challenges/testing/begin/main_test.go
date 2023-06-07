// challenges/testing/begin/main_test.go
package main

import (
	"testing"
)

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{name: "first", input: "abc", want: 3},
		{name: "second", input: "abcd", want: 5},
	}

	letterCounter := letterCounter{identifier: "lettercounter"}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if letterCounter.count(test.input) != test.want {
				t.Errorf("Error in %s, want %d, input is %s", test.name, test.want, test.input)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"first":  {input: "12", want: 2},
		"second": {input: "34", want: 3},
	}

	numberCounter := numberCounter{designation: "numbercounter"}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if got := numberCounter.count(test.input); got != test.want {
				t.Errorf("Error in %s, want %d, input is %s", numberCounter.designation, test.want, test.input)
			}
		})

	}
}

// write a test for symbolCounter.count
