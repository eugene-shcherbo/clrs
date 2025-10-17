package heap

import "cmp"

type Comparator[T any] func(a, b T) int

func DefaultComparator[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

type HeapProperty[T any] struct {
	Satisfies func(a, b T) bool
}

func MinHeapProp[T any](comp Comparator[T]) *HeapProperty[T] {
	return &HeapProperty[T]{func(a, b T) bool { return comp(a, b) <= 0 }}
}

func MaxHeapProp[T any](comp Comparator[T]) *HeapProperty[T] {
	return &HeapProperty[T]{func(a, b T) bool { return comp(a, b) >= 0 }}
}

func BuildHeapInPlace[T any](items []T, prop *HeapProperty[T]) []T {
	for i := len(items)/2 - 1; i >= 0; i-- {
		Heapify(items, i, len(items), prop)
	}
	return items
}

func Heapify[T any](items []T, i int, n int, prop *HeapProperty[T]) []T {
	for {
		idx := i
		left := getLeftIdx(i)
		right := getRightIdx(i)

		if left < n && prop.Satisfies(items[left], items[idx]) {
			idx = left
		}

		if right < n && prop.Satisfies(items[right], items[idx]) {
			idx = right
		}

		if idx == i {
			break
		}

		items[i], items[idx] = items[idx], items[i]
		i = idx
	}

	return items
}

func getLeftIdx(i int) int {
	return 2*i + 1
}

func getRightIdx(i int) int {
	return 2*i + 2
}
