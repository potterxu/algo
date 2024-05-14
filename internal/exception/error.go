package exception

import "errors"

var (
	ErrEmptySlice     = errors.New("empty slice")
	ErrNotFound       = errors.New("target not found")
	ErrSliceNotSorted = errors.New("arr is not sorted")
)
