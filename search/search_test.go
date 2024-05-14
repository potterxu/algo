package search_test

import (
	"testing"

	"github.com/potterxu/algo/search"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6}
	arr2 = []int{1, 2, 2, 4, 5, 6}
)

func TestLowerBound(t *testing.T) {
	arr := arr1
	target := 3
	want := 2
	if got, _ := search.LowerBound(arr, target); got != want {
		t.Errorf("LowerBound(%v, %v) = %v, want %v", arr, target, got, want)
	}

	arr = arr2
	target = 3
	want = 3
	if got, _ := search.LowerBound(arr, target); got != want {
		t.Errorf("LowerBound(%v, %v) = %v, want %v", arr, target, got, want)
	}
}

func TestUpperBound(t *testing.T) {
	arr := arr1
	target := 3
	want := 3
	if got, _ := search.UpperBound(arr, target); got != want {
		t.Errorf("UpperBound(%v, %v) = %v, want %v", arr, target, got, want)
	}

	arr = arr2
	target = 3
	want = 3
	if got, _ := search.UpperBound(arr, target); got != want {
		t.Errorf("UpperBound(%v, %v) = %v, want %v", arr, target, got, want)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := arr1
	target := 3
	want := 2
	if got, _ := search.BinarySearch(arr, target); got != want {
		t.Errorf("BinarySearch(%v, %v) = %v, want %v", arr, target, got, want)
	}

	arr = arr2
	target = 3
	want = -1
	if got, _ := search.BinarySearch(arr, target); got != want {
		t.Errorf("BinarySearch(%v, %v) = %v, want %v", arr, target, got, want)
	}
}

func TestLinearSearch(t *testing.T) {
	arr := arr1
	target := 3
	want := 2
	if got, _ := search.LinearSearch(arr, target); got != want {
		t.Errorf("LinearSearch(%v, %v) = %v, want %v", arr, target, got, want)
	}

	arr = arr2
	target = 3
	want = -1
	if got, _ := search.LinearSearch(arr, target); got != want {
		t.Errorf("LinearSearch(%v, %v) = %v, want %v", arr, target, got, want)
	}
}
