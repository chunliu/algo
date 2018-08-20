package algo

import (
	"testing"
)

func strSliceEqual(a, b []string) bool {
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
func TestTowerOfHanoi(t *testing.T) {
	expected := []string{"from A to C", "from A to B", "from C to B", "from A to C", "from B to A", "from B to C", "from A to C"}
	result := TowerOfHanoi(3)
	if !strSliceEqual(expected, result) {
		t.Errorf("Failed. Expected: %v. Got: %v", expected, result)
	}
}
