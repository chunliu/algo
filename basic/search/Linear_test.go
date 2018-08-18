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
