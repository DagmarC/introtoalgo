package queue

import (
	"fmt"
	"testing"
)

const DEQUEU string = "dequeue" // if error occurs "dequeue" is returned
const DEQUEUI int = 0

func TestEnqueue(t *testing.T) {
	q := NewQueue[float64]()
	q.Enqueue(0.21)
	q.Enqueue(0.33)
	if q.tail != 2 || q.head != 0 {
		t.Errorf("expected head index %d, tail index %d, got head %d, tail %d", 0, 2, q.head, q.tail)
	}
	if q.arr[q.head] != 0.21 || q.arr[q.tail-1] != 0.33 {
		t.Errorf("expected head value %f, last value %f, got head %f, tail %f", 0.21, 0.33, q.arr[q.head], q.arr[q.tail-1])
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(2)
	q.Enqueue(3)
	el, err := q.Dequeue(DEQUEUI)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if q.tail != 1 || q.head != 0 {
		t.Errorf("expected head index %d, tail index %d, got head %d, tail %d", 0, 1, q.head, q.tail)
	}
	if el != 2 {
		t.Errorf("expected dequeued value %d, got %d", 2, el)
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hello")
	q.Enqueue("world")

	el, err := q.Dequeue(DEQUEU)
	if err != nil {
		t.Error("unexpected error", err)
	}
	fmt.Printf(el + " ")

	el, err = q.Dequeue(DEQUEU)
	if err != nil {
		t.Error("unexpected error", err)
	}
	fmt.Printf(el + "\n")

	_, err = q.Dequeue(DEQUEU)
	if err == nil {
		t.Errorf("expected empty queue error when dequeue, got queue head %d", q.head)
	}

}
