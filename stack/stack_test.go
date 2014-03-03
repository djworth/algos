package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := New()

	s.Push("Hello")
	s.Push("World")
	s.Push("This is a")
	s.Push("Stack")

	if s.Empty() {
		t.Error("Expected the stack to not be empty")
	}

	top := s.Pop()

	if top != "Stack" {
		t.Errorf("Expected the top of the stack to equal 'Stack', got %v instead", top)
	}
}

func TestPop(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push("three")
	s.Push("four")

	if s.Empty() {
		t.Error("Expected the stack to not be empty")
	}

	four := s.Pop()
	three := s.Pop()
	two := s.Pop()
	one := s.Pop()

	if four != "four" {
		t.Errorf("Expected the top of the stack to equal 'four', got %v instead", four)
	}

	if three != "three" {
		t.Errorf("Expected the top of the stack to equal 'three', got %v instead", four)
	}

	if two != 2 {
		t.Errorf("Expected the top of the stack to equal 2, got %v instead", two)
	}

	if one != 1 {
		t.Errorf("Expected the top of the stack to equal 'four', got %v instead", one)
	}
}

func TestEmpty(t *testing.T) {
	s := New()

	if !s.Empty() {
		t.Error("Expected the stack to be empty")
	}

	s.Push("one")
	s.Push("two")

	if s.Empty() {
		t.Error("Expected the stack to not be empty")
	}

	s.Pop()
	s.Pop()

	if !s.Empty() {
		t.Error("Expected the stack to be empty")
	}

}

func TestPeek(t *testing.T) {
	s := New()

	if !s.Empty() {
		t.Error("Expected the stack to be empty")
	}

	s.Push("one")
	one := s.Peek()

	if s.Empty() {
		t.Error("Expected the stack to not be empty")
	}

	if one != "one" {
		t.Errorf("Expected the top of the stack to equal 'one', go %v instead", one)
	}

	s.Push("two")
	two := s.Peek()

	if two != "two" {
		t.Errorf("Expected the top of the stack to equal 'two', go %v instead", two)
	}

}
