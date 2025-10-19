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
