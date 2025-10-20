package ch6_test

import (
	"clrs/internal/problems/ch6"
	"clrs/test/assert"
	"testing"
)

func TestKMerge(t *testing.T) {
	lists := [][]int{
		{-1, 5, 23, 112, 154},
		{-20, 100, 115},
		{500, 501},
		{-100, -99, -98, -1, 0},
	}

	want := []int{-100, -99, -98, -20, -1, -1, 0, 5, 23, 100, 112, 115, 154, 500, 501}

	got := ch6.MergeSortedLists(lists)

	assert.SlicesEqual(t, want, got)
}
