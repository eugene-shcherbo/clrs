package sorting

import (
	"clrs/algorithms/slices"
)

func NewQuickSorter[T any](pivotFn func([]T) int) Sorter[T] {
	var sorter Sorter[T]

	sorter = func(items []T, less func(a, b T) bool) []T {
		if len(items) <= 1 {
			return items
		}

		pivotIdx := slices.PartitionInPlace(items, pivotFn(items), func(a, b T) bool { return less(a, b) })
		sorter(items[:pivotIdx], less)
		sorter(items[pivotIdx+1:], less)

		return items
	}

	return sorter
}
