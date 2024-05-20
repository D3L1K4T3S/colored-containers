package src

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 2}, {2, 1}},
			expected: [][]int{{3, 0}, {0, 3}},
		},
		{
			input:    [][]int{{10, 20, 30}, {1, 1, 1}, {0, 0, 1}},
			expected: [][]int{{11, 19, 30}, {0, 2, 1}, {0, 0, 1}},
		},
		{
			input:    [][]int{{10, 20, 30}, {20, 10, 20}, {30, 20, 10}},
			expected: [][]int{{60, 0, 0}, {0, 50, 0}, {0, 0, 60}},
		},
		{
			input:    [][]int{{10, 20, 30}, {20, 10, 20}, {30, 20, 10}},
			expected: [][]int{{60, 0, 0}, {0, 50, 0}, {0, 0, 60}},
		},
	}

	for _, c := range cases {
		solve(c.input)
		if reflect.DeepEqual(c.input, c.expected) == false {
			t.Errorf("result: %v, expected: %v", c.input, c.expected)
		}
	}
}

func TestIsPossible(t *testing.T) {
	cases := []struct {
		input          [][]int
		expectedAnswer string
	}{
		{
			input:          [][]int{{3, 0}, {0, 3}},
			expectedAnswer: "yes",
		},
	}

	for _, c := range cases {
		var answer string
		if isPossible(c.input) {
			answer = "yes"
		} else {
			answer = "no"
		}
		if answer != c.expectedAnswer {
			t.Errorf("result: %v, expected: %v", answer, c.expectedAnswer)
		}
	}
}
