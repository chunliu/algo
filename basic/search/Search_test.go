package algo

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	i := LinearSearch(items, 46)
	if i != 2 {
		t.Errorf("Failed. Expected 2, got %v", i)
	}
}

func TestBinarySearch(t *testing.T) {
	items := []int{1, 3, 5, 6, 10, 22, 34}
	i := BinarySearch(items, 6)
	if i != 3 {
		t.Errorf("Failed. Expected 3, got %v", i)
	}

	i = BinarySearch(items, 50)
	if i != -1 {
		t.Errorf("Failed. Expected -1, got %v", i)
	}
}
