package aoc_utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func WaitForInput() {
	var wait string
	fmt.Scanln(&wait)
}

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxF[T constraints.Ordered](initial T) func(...T) T {
	max := initial
	return func(new ...T) T {
		for _, a := range new {
			if a > max {
				max = a
			}
		}
		return max
	}
}

func MinOf[T constraints.Ordered](vars ...T) (T, int) {
	min := vars[0]
	i := 0

	for j, n := range vars {
		if min > n {
			min, i = n, j
		}
	}

	return min, i
}

func MaxOf[T constraints.Ordered](vars ...T) (T, int) {
	max := vars[0]
	i := 0

	for j, n := range vars {
		if max < n {
			max, i = n, j
		}
	}

	return max, i
}

func Sum[T constraints.Integer | constraints.Float](vars ...T) T {
	var s T
	for _, n := range vars {
		s += n
	}
	return s
}
