package sorting

import (
	"clrs/ds/heap"
)

func HeapsortInPlace[T any](items []T, less func(a, b T) bool) []T {
	prop := heap.MaxHeapProp(func(a, b T) int {
		if less(a, b) {
			return -1
		}

		return 1
	})

	heap.BuildHeapInPlace(items, prop)

	for i := len(items) - 1; i >= 0; i-- {
		items[i], items[0] = items[0], items[i]
		heap.Heapify(items, 0, i, prop)
	}

	return items
}
