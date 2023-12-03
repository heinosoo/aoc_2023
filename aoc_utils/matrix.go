package aoc_utils

import "strings"

type Matrix[T any] [][]T

func InitializeMatrix[T any](value T, M, N int) (matrix Matrix[T]) {
	for m := 0; m < M; m++ {
		matrix = append(matrix, make([]T, N))
		for n := 0; n < N; n++ {
			matrix[m][n] = value
		}
	}
	return
}

// CreateMatrixFromInputChannel creates a 2D matrix from a channel of strings, where each string represents a row.
// The whole channel will be consumed.
func CreateMatrixFromInputChannel(lines <-chan string) (matrix Matrix[string]) {
	for line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	return
}

func (matrix Matrix[T]) Print() {
	M, N := len(matrix), len((matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			print(matrix[m][n])
		}
		print("\n")
	}
}

func (matrix Matrix[T]) ForEach(function func(T)) {
	M, N := len(matrix), len((matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			function((matrix)[m][n])
		}
	}
}

func (matrix Matrix[T]) Apply(function func(T) T) {
	M, N := len(matrix), len((matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			matrix[m][n] = function(matrix[m][n])
		}
	}
}

func (matrix Matrix[T]) Pad(paddingValue T) (matrixPadded Matrix[T]) {
	paddedWidth := len(matrix[0]) + 2

	paddingRow := make([]T, paddedWidth)
	for i := range paddingRow {
		paddingRow[i] = paddingValue
	}

	matrixPadded = append(matrixPadded, paddingRow)
	for _, row := range matrix {
		paddedRow := make([]T, 0, paddedWidth)
		paddedRow = append(paddedRow, paddingValue)
		paddedRow = append(paddedRow, row...)
		paddedRow = append(paddedRow, paddingValue)
		matrixPadded = append(matrixPadded, paddedRow)
	}
	matrixPadded = append(matrixPadded, paddingRow)

	return
}
