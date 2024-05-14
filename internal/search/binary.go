package search

import (
	"cmp"
	"sort"

	"github.com/potterxu/algo/internal/exception"
)

func check[T cmp.Ordered](arr []T) error {
	if len(arr) == 0 {
		return exception.ErrEmptySlice
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < (arr[i-1]) {
			return exception.ErrSliceNotSorted
		}
	}
	return nil
}

// use binary search to find the first element that is not less than target
func LowerBound[T cmp.Ordered](arr []T, target T) (int, error) {
	if err := check(arr); err != nil {
		return -1, err
	}
	index, _ := sort.Find(len(arr), func(i int) int {
		return cmp.Compare(target, arr[i])
	})
	if index == len(arr) {
		return -1, exception.ErrNotFound
	}
	return index, nil
}

// use binary search to find the first element that is greater than target
func UpperBound[T cmp.Ordered](arr []T, target T) (int, error) {
	if err := check(arr); err != nil {
		return -1, err
	}
	index, _ := sort.Find(len(arr), func(i int) int {
		if arr[i] > target {
			return -1
		} else {
			return 1
		}
	})
	if index == len(arr) {
		return -1, exception.ErrNotFound
	}
	return index, nil
}
