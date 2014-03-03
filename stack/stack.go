package stack

import (
	"container/list"
)

// Stack interface supporting the normal stack functions.  The implementation uses a linked list
type Stack interface {
	// Push puts an item on the top of the stack
	Push(item interface{})
	// Pop removes an item from the top of the stack
	Pop() interface{}
	// Peek looks at the item at the top of the stack but does not remove the item
	Peek() interface{}
	// Empty returns if the stack has items or not
	Empty() bool
}

type stack struct {
	l *list.List
}

func New() Stack {
	return &stack{l: list.New()}
}

func (s *stack) Push(item interface{}) {
	s.l.PushFront(item)
}

func (s *stack) Pop() interface{} {

	if s.Empty() {
		return nil
	}

	e := s.l.Front()
	s.l.Remove(e)
	return e.Value
}

func (s *stack) Peek() interface{} {

	if s.Empty() {
		return nil
	}

	e := s.l.Front()
	return e.Value
}

func (s *stack) Empty() bool {
	return s.l.Len() == 0
}
