package random_test

import (
	"clrs/algorithms/random"
	"slices"
	"testing"
)

func TestPermuteShouldProvideAllElements(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	permuted := random.Permute(items)

	if len(permuted) != len(items) {
		t.Errorf("Expected length %d, got %d", len(items), len(permuted))
	}

	for _, v := range items {
		if slices.Index(permuted, v) == -1 {
			t.Errorf("Element %d from original slice not found in permuted slice", v)
		}
	}
}
