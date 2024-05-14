package search

import (
	"github.com/potterxu/algo/internal/search"
)

// find minimum index i such that arr[i] == target
func LinearSearch[T comparable](arr []T, target T) (int, error) {
	return search.LinearSearch(arr, target)
}
