package sorting

import (
	"clrs/test/assert"
	"testing"
)

func TestInPlaceSorting(t *testing.T) {
	algorithms := []Sorter[int, int]{HeapsortInPlace[int, int]}

	want := []int{-29, -1, 0, 1, 3, 8, 8, 10, 10, 35, 54, 120}
	for _, algo := range algorithms {
		input := []int{8, 10, 54, 1, 0, -29, 35, 120, 3, 8, 10, -1}

		got := algo(input, func(a int) int { return a })

		assert.SlicesEqual(t, want, got)
		assert.SlicesEqual(t, want, input)
	}
}
