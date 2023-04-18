package queue

import (
	"errors"
)

type Elems interface {
	~int | ~string | ~float64
}
type Queue[T Elems] struct {
	tail int // index of the location of new incoming element
	head int // index of the first element enqueued, empty queue = -1
	arr  []T
}

func NewQueue[T Elems]() *Queue[T] {
	return &Queue[T]{tail: 0, head: -1, arr: make([]T, 0)}
}

// Enqueue the element and sets the tail and head accordingly
func (q *Queue[T]) Enqueue(el T) {
	if q.head == -1 {
		q.head++ // first element being inserted, q is not empty
	}
	q.arr = append(q.arr, el)
	q.tail++
}

// Dequeue removes the first element (head element) from the array,
// if the queue is empty parameter e with error is returned
func (q *Queue[T]) Dequeue(e T) (T, error) {
	if q.tail == 0 || q.head == -1 {
		return e, errors.New("empty queue")
	}
	el := q.arr[q.head]
	q.arr = q.arr[q.head+1 : q.tail]
	q.tail-- // Note: in arr head is always 0 if there exists at least 1 element, otherwise head is -1

	return el, nil
}
