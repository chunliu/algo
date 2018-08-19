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
