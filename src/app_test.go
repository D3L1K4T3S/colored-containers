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
	}

	for _, c := range cases {
		solve(c.input)
		if reflect.DeepEqual(c.input, c.expected) == false {
			t.Errorf("result: %v, expected: %v", c.input, c.expected)
		}
	}
}

//func TestIsPossible(t *testing.T) {
//
//}
