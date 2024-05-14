package search

import (
	"github.com/potterxu/algo/internal/exception"
)

// find minimum index i such that arr[i] == target
func LinearSearch[T comparable](arr []T, target T) (int, error) {
	if len(arr) == 0 {
		return -1, exception.ErrEmptySlice
	}
	for i, v := range arr {
		if v == target {
			return i, nil
		}
	}
	return -1, exception.ErrNotFound
}
