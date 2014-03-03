package queue

import (
	"testing"
)

func TestPushAndRemove(t *testing.T) {

	q := New()

	q.Add("hello")
	q.Add("world")

	if q.Empty() {
		t.Error("q should not be empty.")
	}

	hello := q.Remove()

	if hello != "hello" {
		t.Errorf("Expected variable hello to equal 'hello', Got %v instead", hello)
	}

	q.Add("danny")

	world := q.Remove()
	if world != "world" {
		t.Errorf("Expected variable world to equal 'world', Got %v instead", world)
	}

	danny := q.Remove()
	if danny != "danny" {
		t.Errorf("Expected variable danny to equal 'danny', Got %v instead", danny)
	}

}

func TestPeek(t *testing.T) {

	q := New()

	q.Add("hello")
	q.Add("world")

	if q.Empty() {
		t.Error("q should not be empty.")
	}

	hello := q.Peek()

	if hello != "hello" {
		t.Errorf("Expected variable hello to equal 'hello', Got %v instead", hello)
	}

	q.Add("danny")

	hello = q.Peek()
	if hello != "hello" {
		t.Errorf("Expected variable hello to equal 'hello', Got %v instead", hello)
	}
}

func TestEmpty(t *testing.T) {
	q := New()

	if !q.Empty() {
		t.Error("q should be empty.")
	}

	q.Add("hello")
	q.Add("world")

	if q.Empty() {
		t.Error("q should not be empty.")
	}
}
