package sorting

import (
	"clrs/ds/heap"
	"cmp"
)

func HeapsortInPlace[T any, K cmp.Ordered](items []T, key func(i T) K) []T {
	comparator := func(a, b T) int { return heap.DefaultComparator(key(a), key(b)) }
	prop := heap.MaxHeapProp(comparator)

	heap.BuildHeapInPlace(items, prop)

	for i := len(items) - 1; i >= 0; i-- {
		items[i], items[0] = items[0], items[i]
		heap.Heapify(items, 0, i, prop)
	}

	return items
}
