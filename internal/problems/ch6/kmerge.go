package ch6

import "clrs/ds/heap"

func MergeSortedLists(lists [][]int) []int {
	type elInfo struct {
		idx     int
		listIdx int
	}

	prop := heap.MinHeapProp(func(e *elInfo) int { return lists[e.listIdx][e.idx] })
	minimums := heap.NewHeap([]*elInfo{}, prop)

	for i := range lists {
		minimums.Add(&elInfo{0, i})
	}

	sorted := make([]int, 0)

	for minimums.Len() > 0 {
		minInfo, _ := minimums.Pop()
		list := lists[minInfo.listIdx]

		sorted = append(sorted, list[minInfo.idx])

		if minInfo.idx < len(list)-1 {
			minimums.Add(&elInfo{minInfo.idx + 1, minInfo.listIdx})
		}
	}

	return sorted
}
