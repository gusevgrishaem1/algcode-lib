package algcode_lib

import "sync"

// syncStack is thread-safe stack implementation that supports concurrent access
type syncStack struct {
	stack *stack
	mu    *sync.RWMutex
}

// NewStackSync return a pointer to a new syncStack
func NewStackSync(cap int) *syncStack {
	return &syncStack{
		stack: NewStack(cap),
		mu:    &sync.RWMutex{},
	}
}

// Push an element onto the top of the stack
func (s *syncStack) Push(v any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stack.Push(v)
}

// Size returns the current size of the stack
func (s *syncStack) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.stack.Size()
}

// Pop removes and returns the top element from stack in LIFO order
func (s *syncStack) Pop() any {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stack.Pop()
}
