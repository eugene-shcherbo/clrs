package assert

import (
	"slices"
	"testing"
)

func Equals[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Errorf("Expected is %v, Actual is %v", a, b)
	}
}

func SlicesEqual[T comparable](t *testing.T, a, b []T) {
	if !slices.Equal(a, b) {
		t.Errorf("Expected is %v, Actual is %v", a, b)
	}
}

func True(t *testing.T, value bool) {
	if !value {
		t.Error("Value should be true")
	}
}

func False(t *testing.T, value bool) {
	if value {
		t.Error("Value should be false")
	}
}
