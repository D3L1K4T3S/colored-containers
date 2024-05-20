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
		{
			input:    [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expected: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			input:    [][]int{{5, 3, 2}, {3, 5, 4}, {2, 4, 0}},
			expected: [][]int{{10, 0, 0}, {0, 12, 0}, {0, 0, 6}},
		},
		{
			input:    [][]int{{100000, 200000, 300000, 400000}, {200000, 100000, 400000, 300000}, {300000, 400000, 500000, 400000}, {400000, 300000, 400000, 500000}},
			expected: [][]int{{1000000, 0, 0, 0}, {0, 1000000, 0, 0}, {0, 0, 1600000, 0}, {0, 0, 0, 1600000}},
		},
		{
			input:    [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			expected: [][]int{{3, 0, 0}, {0, 3, 0}, {0, 0, 3}},
		},
		{
			input:    [][]int{{100000000, 1}, {1, 100000000}},
			expected: [][]int{{100000001, 0}, {0, 100000001}},
		},
		{
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			input:    [][]int{{1}},
			expected: [][]int{{1}},
		},
		{
			input:    [][]int{{0, 0, 3}, {0, 3, 0}, {3, 0, 0}},
			expected: [][]int{{3, 0, 0}, {0, 3, 0}, {0, 0, 3}},
		},
		{},
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
		{
			input:          [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expectedAnswer: "yes",
		},
		{
			input:          [][]int{{0}},
			expectedAnswer: "yes",
		},
		{
			input:          [][]int{{1, 2}, {1, 2}},
			expectedAnswer: "no",
		},
		{
			input:          [][]int{{1}},
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
