package heap

import (
	"cmp"
	"errors"
)

func DefaultKey[T cmp.Ordered](a T) T {
	return a
}

type HeapProperty[T any, K cmp.Ordered] struct {
	Satisfies func(a, b T) bool
	keyOf     func(T) K
}

func MinHeapProp[T any, K cmp.Ordered](keyOf func(T) K) *HeapProperty[T, K] {
	return &HeapProperty[T, K]{
		func(a, b T) bool { return keyOf(a) <= keyOf(b) },
		keyOf,
	}
}

func MaxHeapProp[T any, K cmp.Ordered](keyOf func(T) K) *HeapProperty[T, K] {
	return &HeapProperty[T, K]{
		func(a, b T) bool { return keyOf(a) >= keyOf(b) },
		keyOf,
	}
}

type BinaryHeap[T any, K cmp.Ordered] struct {
	prop  *HeapProperty[T, K]
	items []T
}

func NewHeap[T any, K cmp.Ordered](items []T, prop *HeapProperty[T, K]) *BinaryHeap[T, K] {
	itemsCopy := make([]T, len(items))
	copy(itemsCopy, items)

	return &BinaryHeap[T, K]{
		prop,
		BuildHeapInPlace(itemsCopy, prop),
	}
}

func (heap *BinaryHeap[T, K]) Len() int {
	return len(heap.items)
}

func (heap *BinaryHeap[T, K]) Peek() (T, error) {
	if len(heap.items) == 0 {
		var zero T
		return zero, errors.New("heap underflow")
	}

	return heap.items[0], nil
}

func (heap *BinaryHeap[T, K]) Pop() (T, error) {
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

func (heap *BinaryHeap[T, K]) Add(item T) {
	heap.items = append(heap.items, item)
	prop, items := heap.prop, heap.items

	for i := len(items) - 1; i > 0; {
		parentIdx := getParentIdx(i)

		if prop.Satisfies(items[parentIdx], items[i]) {
			return
		}

		items[i], items[parentIdx] = items[parentIdx], items[i]
		i = parentIdx
	}
}

func BuildHeapInPlace[T any, K cmp.Ordered](items []T, prop *HeapProperty[T, K]) []T {
	for i := len(items)/2 - 1; i >= 0; i-- {
		Heapify(items, i, len(items), prop)
	}
	return items
}

func Heapify[T any, K cmp.Ordered](items []T, i int, n int, prop *HeapProperty[T, K]) []T {
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
