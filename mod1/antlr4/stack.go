package antlr4

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		pre   *node
	}
)

func NewStack() *stack {
	return &stack{nil, 0}
}

func (s *stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.pre
	s.length--
	return n.value
}

func (s *stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
