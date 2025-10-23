package sorting

import (
	"clrs/algorithms/random"
	"clrs/test/assert"
	"testing"
)

func TestInPlaceSorting(t *testing.T) {
	algorithms := []Sorter[int]{
		HeapsortInPlace[int],
		NewQuickSorter(func(items []int) int { return len(items) - 1 }),
		NewQuickSorter(func(items []int) int { return random.RandRange(0, len(items)) }),
	}

	want := []int{-29, -1, 0, 1, 3, 8, 8, 10, 10, 35, 54, 120}
	for _, algo := range algorithms {
		input := []int{8, 10, 54, 1, 0, -29, 35, 120, 3, 8, 10, -1}

		got := algo(input, func(a, b int) bool { return a < b })

		assert.SlicesEqual(t, want, got)
		assert.SlicesEqual(t, want, input)
	}
}
