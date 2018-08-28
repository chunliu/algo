package tree

type PriorityQueue struct {
	data []int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq *PriorityQueue) Enqueue(data int) {
	pq.data = append(pq.data, data)

	index := len(pq.data) - 1
	for index > 0 {
		parent := (index - 1) / 2
		if pq.data[parent] >= pq.data[index] {
			break
		}
		pq.data[parent], pq.data[index] = pq.data[index], pq.data[parent]
		index = parent
	}
}

func (pq *PriorityQueue) Dequeue() int {
	if len(pq.data) == 0 {
		return 0
	}

	ret := pq.data[0]
	end := len(pq.data) - 1
	pq.data[0], pq.data[end] = pq.data[end], pq.data[0]
	pq.data = pq.data[:end]
	if len(pq.data) == 0 {
		return ret
	}

	index := 0
	for true {
		child1 := 2*index + 1
		child2 := 2*index + 2
		// If child index is bigger than the count, there is no more child.
		if child1 >= len(pq.data) {
			child1 = index
		}
		if child2 >= len(pq.data) {
			child2 = index
		}
		// If item index is bigger than both children, it is already a heap.
		if pq.data[index] >= pq.data[child1] && pq.data[index] >= pq.data[child2] {
			break
		}
		// Otherwise, swap the item index with the bigger child. Then repeat to push it further down.
		if pq.data[child1] > pq.data[child2] {
			pq.data[child1], pq.data[index] = pq.data[index], pq.data[child1]
			index = child1
		} else {
			pq.data[child2], pq.data[index] = pq.data[index], pq.data[child2]
			index = child2
		}
	}

	return ret
}

func (pq *PriorityQueue) Len() int {
	return len(pq.data)
}
