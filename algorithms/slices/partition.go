package slices

func PartitionInPlace[T any](items []T, pivotIdx int, less func(a, b T) bool) int {
	n := len(items)
	pivot := items[pivotIdx]
	items[pivotIdx], items[n-1] = items[n-1], items[pivotIdx]

	lowIdx := -1
	for i := range n - 1 {
		if less(items[i], pivot) {
			lowIdx++
			items[lowIdx], items[i] = items[i], items[lowIdx]
		}
	}

	items[lowIdx+1], items[n-1] = items[n-1], items[lowIdx+1]

	return lowIdx + 1
}
