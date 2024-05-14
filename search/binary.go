package search

import (
	"cmp"

	"github.com/potterxu/algo/internal/exception"
	"github.com/potterxu/algo/internal/search"
)

func BinarySearch[T cmp.Ordered](arr []T, target T) (int, error) {
	index, err := search.LowerBound(arr, target)
	if err != nil {
		return -1, err
	}
	if arr[index] != target {
		return -1, exception.ErrNotFound
	}
	return index, nil
}

// use binary search to find the first element that is not less than target
func LowerBound[T cmp.Ordered](arr []T, target T) (int, error) {
	return search.LowerBound(arr, target)
}

// use binary search to find the first element that is greater than target
func UpperBound[T cmp.Ordered](arr []T, target T) (int, error) {
	return search.UpperBound(arr, target)
}
