package sorting

import "cmp"

type Sorter[T any, K cmp.Ordered] func(items []T, key func(a T) K) []T
