package queues

import (
	"fmt"
	"github.com/246859/containers"
	"github.com/246859/containers/internal"
)

var _ Queue[any] = (*ArrayQueue[any])(nil)

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{s: make([]T, 0, capacity)}
}

// ArrayQueue implements by slice, one of the most easily implementation of Queue
type ArrayQueue[T any] struct {
	s []T
}

func (queue *ArrayQueue[T]) Push(es ...T) {
	queue.s = append(queue.s, es...)
}

func (queue *ArrayQueue[T]) Peek() (_ T, _ bool) {
	if queue.Size() == 0 {
		return
	}
	return queue.s[0], true
}

func (queue *ArrayQueue[T]) Pop() (_ T, _ bool) {
	if queue.Size() == 0 {
		return
	}
	peek := queue.s[0]
	queue.s = queue.s[1:]
	return peek, true
}

func (queue *ArrayQueue[T]) Iterator() containers.IndexIterator[T] {
	return internal.NewSliceIterator(queue.s)
}

func (queue *ArrayQueue[T]) Values() []T {
	var vals []T
	it := queue.Iterator()
	for it.Next() {
		vals = append(vals, it.Value())
	}
	return vals
}

func (queue *ArrayQueue[T]) Size() int {
	return len(queue.s)
}

func (queue *ArrayQueue[T]) Clear() {
	queue.s = queue.s[:0:0]
}

func (queue *ArrayQueue[T]) String() string {
	return fmt.Sprintf("%+v", queue.s)
}
