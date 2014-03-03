package queue

import (
	"github.com/djworth/algos/stack"
)

// Queue is the interface of a queue implementation using two stacks
type Queue interface {
	// Add puts an item on the queue
	Add(item interface{})
	// Remove takes the first time off of the queue
	Remove() interface{}
	// Peek returns the first on the queue
	Peek() interface{}
	// Empty returns true or false if the queue contains values
	Empty() bool
}

type queue struct {
	s1 stack.Stack
	s2 stack.Stack
}

func New() *queue {
	return &queue{s1: stack.New(), s2: stack.New()}
}

func (q *queue) Add(item interface{}) {
	if !q.s2.Empty() {
		q.flush(q.s2, q.s1)
	}

	q.s1.Push(item)

}

func (q *queue) Remove() interface{} {

	if q.Empty() {
		return nil
	}

	q.flush(q.s1, q.s2)

	return q.s2.Pop()
}

func (q *queue) Peek() interface{} {
	if q.Empty() {
		return nil
	}

	q.flush(q.s1, q.s2)

	return q.s2.Peek()
}

func (q *queue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
}

func (q *queue) flush(from, to stack.Stack) {
	for {
		item := from.Pop()
		if item == nil {
			break
		}
		to.Push(item)
	}
}
