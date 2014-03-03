package stack

import (
	"container/list"
)

type Stack interface {
	Push(item interface{})
	Pop() interface{}
	Peek() interface{}
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
