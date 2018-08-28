package tree

import "testing"

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

func TestHeapSort(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	expected := []int{45, 46, 58, 78, 86, 95, 99, 251, 320}
	sorted := HeapSort(items)
	if !sliceEqual(expected, sorted) {
		t.Errorf("Failed. Expected: %v; Got: %v", expected, sorted)
	}
}

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue()
	q.Enqueue(2)
	q.Enqueue(4)
	q.Enqueue(6)
	q.Enqueue(8)

	v := q.Dequeue()
	if v != 8 {
		t.Errorf("Failed. Expected: 8. Get: %v", v)
	}
	q.Dequeue()
	v = q.Dequeue()
	if v != 4 {
		t.Errorf("Failed. Expected: 4. Get: %v", v)
	}
	q.Dequeue()
}
