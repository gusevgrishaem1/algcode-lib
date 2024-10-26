package algcode_lib

// stack - a data structure that stores elements in LIFO (Last In, First Out) order.
// A stack uses LIFO (Last In, First Out) ordering to manage data storage and retrieval.
type stack struct {
	elems []any
}

// NewStack return a pointer to a new stack
func NewStack(cap int) *stack {
	return &stack{elems: make([]any, 0, cap)}
}

// Push an element onto the top of the stack
func (s *stack) Push(v any) {
	s.elems = append(s.elems, v)
}

// Size returns the current size of the stack
func (s *stack) Size() int {
	return len(s.elems)
}

// Pop removes and returns the top element from stack in LIFO order
func (s *stack) Pop() any {
	if len(s.elems) == 0 {
		return nil
	}

	poped := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return poped
}
