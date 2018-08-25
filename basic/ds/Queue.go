package algo

type Queue struct {
	nodes []*node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.nodes = append(q.nodes, &node{v})
}

func (q *Queue) Dequeue() interface{} {
	if len(q.nodes) == 0 {
		return nil
	}

	v := q.nodes[0]
	q.nodes = q.nodes[1:]

	return v.value
}

func (q *Queue) Len() int {
	return len(q.nodes)
}
