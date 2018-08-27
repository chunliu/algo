package ds

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	x := q.Dequeue()
	if x != nil {
		t.Error("Failed.")
	}

	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	if q.Len() != 5 {
		t.Errorf("Failed. Expected: 5, Get: %v", q.Len())
	}

	q.Dequeue()
	q.Dequeue()
	v := q.Dequeue().(int)
	if v != 3 {
		t.Errorf("Failed. Expected: 1, Get: %v", v)
	}
}
