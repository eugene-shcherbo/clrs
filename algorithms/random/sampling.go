package random

import (
	"fmt"
	"maps"
	"slices"
)

func Sample[T any](items []T, m int, key func(T) int) ([]T, error) {
	if m < 0 || m > len(items) {
		return nil, fmt.Errorf("sample size %v is bigger than input size %v", m, len(items))
	}

	result := map[int]T{}

	for i := len(items) - m; i < len(items); i++ {
		value := items[RandRange(0, i)]
		valueKey := key(value)

		if _, ok := result[valueKey]; !ok {
			result[valueKey] = value
		} else {
			result[key(items[i])] = items[i]
		}
	}

	return slices.Collect(maps.Values(result)), nil
}
