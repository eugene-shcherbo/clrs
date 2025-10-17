package random

func Permute[T any](items []T) []T {
	for i := range len(items) {
		j := RandRange(i, len(items))
		items[i], items[j] = items[j], items[i]
	}
	return items
}
