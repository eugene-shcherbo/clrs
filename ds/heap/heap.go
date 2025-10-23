package heap

import (
	"cmp"
	"errors"
)

func DefaultComparer[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

type HeapProperty[T any] struct {
	Satisfies func(a, b T) bool
}

func MinHeapProp[T any](comparer func(a, b T) int) *HeapProperty[T] {
	return &HeapProperty[T]{func(a, b T) bool { return comparer(a, b) < 0 }}
}

func MaxHeapProp[T any](comparer func(a, b T) int) *HeapProperty[T] {
	return &HeapProperty[T]{func(a, b T) bool { return comparer(a, b) > 0 }}
}

type BinaryHeap[T any] struct {
	prop  *HeapProperty[T]
	items []T
}

func NewHeap[T any](items []T, prop *HeapProperty[T]) *BinaryHeap[T] {
	itemsCopy := make([]T, len(items))
	copy(itemsCopy, items)

	return &BinaryHeap[T]{
		prop,
		BuildHeapInPlace(itemsCopy, prop),
	}
}

func (heap *BinaryHeap[T]) Len() int {
	return len(heap.items)
}

func (heap *BinaryHeap[T]) Peek() (T, error) {
	if len(heap.items) == 0 {
		var zero T
		return zero, errors.New("heap underflow")
	}

	return heap.items[0], nil
}

func (heap *BinaryHeap[T]) Pop() (T, error) {
	val, err := heap.Peek()

	if err != nil {
		return val, err
	}

	n := len(heap.items)

	heap.items[0], heap.items[n-1] = heap.items[n-1], heap.items[0]
	heap.items = heap.items[:n-1]

	Heapify(heap.items, 0, len(heap.items), heap.prop)

	return val, nil
}

func (heap *BinaryHeap[T]) Add(item T) {
	heap.items = append(heap.items, item)
	prop, items := heap.prop, heap.items

	i := len(items) - 1
	j := getParentIdx(i)

	for i > 0 && !prop.Satisfies(items[j], item) {
		items[i] = items[j]
		i = j
		j = getParentIdx(i)
	}

	items[i] = item
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

func getParentIdx(i int) int {
	return (i - 1) / 2
}
