package ch5

import "clrs/algorithms/random"

func RandomSearch[T comparable](items []T, value T) int {
	checked := make([]bool, len(items))
	checkedNum := 0

	for checkedNum < len(items) {
		i := random.RandRange(0, len(items))

		if items[i] == value {
			return i
		}

		if !checked[i] {
			checkedNum++
			checked[i] = true
		}
	}

	return -1
}
