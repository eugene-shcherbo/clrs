package random_test

import (
	"clrs/algorithms/random"
	"clrs/test/assert"
	"slices"
	"testing"
)

func TestSubsetOfInput(t *testing.T) {
	original := []int{10, 11, 12, 9, 20, 21}

	sample, err := random.Sample(original, 4, func(i int) int { return i })

	assert.True(t, err == nil)
	assert.Equals(t, 4, len(sample))
	assert.True(t, isEveryItemIn(original, sample))
}

func TestSampleBounds(t *testing.T) {
	input := []int{10, 11, 12}
	testCases := []int{-1, 4, 20}

	for _, c := range testCases {
		sample, err := random.Sample(input, c, func(i int) int { return i })

		assert.True(t, sample == nil)
		assert.True(t, err != nil)
	}
}

func isEveryItemIn(items []int, subset []int) bool {
	for _, v := range subset {
		if slices.Index(items, v) == -1 {
			return false
		}
	}
	return true
}
