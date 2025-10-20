package heap_test

import (
	"clrs/ds/heap"
	"clrs/test/assert"
	"testing"
)

func TestReadForEmpty(t *testing.T) {
	operations := []func(*heap.BinaryHeap[int, int]) (int, error){
		func(heap *heap.BinaryHeap[int, int]) (int, error) { return heap.Peek() },
		func(heap *heap.BinaryHeap[int, int]) (int, error) { return heap.Pop() },
	}

	for _, op := range operations {
		val, err := op(heap.NewHeap(nil, heap.MaxHeapProp(heap.DefaultKey[int])))

		assert.Equals(t, 0, val)
		assert.True(t, err != nil)
	}
}

func TestPeek(t *testing.T) {
	testCases := []struct {
		prop *heap.HeapProperty[int, int]
		want int
	}{
		{heap.MinHeapProp(heap.DefaultKey[int]), -3},
		{heap.MaxHeapProp(heap.DefaultKey[int]), 128},
	}

	for _, c := range testCases {
		heap := heap.NewHeap([]int{4, 8, 1, 10, -3, 11, 128, 2, 2, 0, 5}, c.prop)

		got, err := heap.Peek()

		assert.Equals(t, nil, err)
		assert.Equals(t, c.want, got)
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		prop *heap.HeapProperty[int, int]
		want []int
	}{
		{heap.MinHeapProp(heap.DefaultKey[int]), []int{-3, 0, 1, 2, 2, 4, 5, 8}},
		{heap.MaxHeapProp(heap.DefaultKey[int]), []int{8, 5, 4, 2, 2, 1, 0, -3}},
	}

	for _, c := range testCases {
		heap := heap.NewHeap([]int{4, 8, 1, -3, 2, 2, 0, 5}, c.prop)

		got := make([]int, 0, heap.Len())
		for heap.Len() > 0 {
			item, _ := heap.Pop()
			got = append(got, item)
		}

		assert.Equals(t, 8, len(got))
		assert.SlicesEqual(t, c.want, got)
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		prop      *heap.HeapProperty[int, int]
		wantPeeks []int
	}{
		{heap.MinHeapProp(heap.DefaultKey[int]), []int{4, 4, 1, 1, 0, 0, -3, -3}},
		{heap.MaxHeapProp(heap.DefaultKey[int]), []int{4, 8, 8, 8, 8, 12, 12, 12}},
	}

	toAdd := []int{4, 8, 1, 2, 0, 12, -3, 5}
	for _, c := range testCases {
		heap := heap.NewHeap([]int{}, c.prop)
		for i, item := range toAdd {
			heap.Add(item)

			peek, _ := heap.Peek()
			assert.Equals(t, c.wantPeeks[i], peek)
			assert.Equals(t, i+1, heap.Len())
		}
	}
}
