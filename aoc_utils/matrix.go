package aoc_utils

func OperateOnMatrix[T any](matrix *[][]T, function func(T)) {
	M, N := len(*matrix), len((*matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			function((*matrix)[m][n])
		}
	}
}

func TransformMatrix[T any](matrix *[][]T, function func(T) T) {
	M, N := len(*matrix), len((*matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			(*matrix)[m][n] = function((*matrix)[m][n])
		}
	}
}

func InitializeMatrix[T any](value T, M, N int) (matrix [][]T) {
	for m := 0; m < M; m++ {
		matrix = append(matrix, make([]T, N))
		for n := 0; n < N; n++ {
			matrix[m][n] = value
		}
	}
	return
}
