package src

import (
	"bufio"
	"fmt"
)

func App(in *bufio.Reader, out *bufio.Writer) error {
	matrix := readMatrix(in)
	if matrix == nil {
		return fmt.Errorf("incorrect matrix input")
	}
	solve(matrix)
	outputMatrix(matrix)
	if isPossible(matrix) {
		_, err := fmt.Fprintln(out, "yes")
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintln(out, "no")
		if err != nil {
			return err
		}
	}
	return nil
}

func isPossible(matrix [][]int) bool {
	result := true
	for i := 0; i < len(matrix) && result; i++ {
		for j := 0; j < len(matrix[i]) && result; j++ {
			if i != j {
				if matrix[i][j] != 0 {
					result = false
				}
			}
		}
	}

	return result
}

func solve(matrix [][]int) {
	var trans int
	for row1 := 0; row1 < len(matrix); row1++ {
		for row2 := 0; row2 < len(matrix); row2++ {
			trans = minInt(matrix[row2][row1], matrix[row1][row2])
			matrix[row1][row1] += trans
			matrix[row2][row1] -= trans
			matrix[row2][row2] += trans
			matrix[row1][row2] -= trans
		}
	}
}

func readMatrix(in *bufio.Reader) [][]int {
	var n int
	var matrix [][]int
	_, err := fmt.Fscan(in, &n)
	if err != nil || n < 1 {
		return nil
	}

	matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, err = fmt.Fscan(in, &matrix[i][j])
			if err != nil {
				return nil
			}
		}
	}
	return matrix
}

func minInt(a, b int) int {
	var result int
	if a < b {
		result = a
	} else {
		result = b
	}
	return result
}

func outputMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)-1; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println(matrix[i][len(matrix[i])-1])
	}
}
