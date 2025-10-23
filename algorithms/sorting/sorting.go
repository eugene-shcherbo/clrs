package sorting

type Sorter[T any] func(items []T, less func(a, b T) bool) []T
