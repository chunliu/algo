package algo

type node struct {
	value interface{}
}

type Stack struct {
	nodes []*node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) {
	s.nodes = append(s.nodes, &node{v})
}

func (s *Stack) Pop() interface{} {
	if len(s.nodes) == 0 {
		return nil
	}
	top := len(s.nodes) - 1
	n := s.nodes[top]
	s.nodes = s.nodes[0:top]

	return n.value
}

func (s *Stack) Peek() interface{} {
	if len(s.nodes) == 0 {
		return nil
	}
	top := len(s.nodes) - 1
	n := s.nodes[top]

	return n.value
}

func (s *Stack) Len() int {
	return len(s.nodes)
}
