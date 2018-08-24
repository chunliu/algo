package algo

import (
	"testing"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSelectionSort(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := SelectionSort(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestSelectionSort2(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := SelectionSort2(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestQuickSort(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := QuickSort(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestQuickSort2(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := QuickSort2(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestInsertionSort(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := InsertionSort(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestMergeSort(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := MergeSort(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func EqualIntervals(i1, i2 []*Interval) bool {
	if len(i1) != len(i2) {
		return false
	}

	for i, v := range i1 {
		if v.Start != i2[i].Start || v.End != i2[i].End {
			return false
		}
	}

	return true
}

func TestMergeIntervals(t *testing.T) {
	intervals := []*Interval{
		&Interval{2, 3},
		&Interval{4, 5},
		&Interval{6, 7},
		&Interval{8, 9},
		&Interval{1, 10},
	}
	expected := []*Interval{
		&Interval{1, 10},
	}

	r := MergeIntervals(intervals)
	if !EqualIntervals(expected, r) {
		t.Errorf("Failed. Expected: %v, Get: %v", expected, r)
	}
}
