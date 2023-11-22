package queues

import (
	"github.com/246859/containers"
	"github.com/246859/containers/heaps"
)

var _ Queue[any] = (*PriorityQueue[any])(nil)

// NewPriorityQueue defaults use binary heap to implement
func NewPriorityQueue[T any](capacity int, compare containers.Compare[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: heaps.NewBinaryHeap[T](capacity, compare),
	}
}

// NewPriorityQueueWith returns a PriorityQueue implemented by the given custom heap
func NewPriorityQueueWith[T any](heap heaps.Heap[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: heap,
	}
}

// PriorityQueue implements by heap
type PriorityQueue[T any] struct {
	heap heaps.Heap[T]
}

func (p *PriorityQueue[T]) Push(es ...T) {
	p.heap.Push(es...)
}

func (p *PriorityQueue[T]) Peek() (_ T, _ bool) {
	return p.heap.Peek()
}

func (p *PriorityQueue[T]) Pop() (_ T, _ bool) {
	return p.heap.Pop()
}

func (p *PriorityQueue[T]) Iterator(reverse bool) containers.IndexIterator[T] {
	return p.heap.Iterator(reverse)
}

func (p *PriorityQueue[T]) Values() []T {
	return p.heap.Values()
}

func (p *PriorityQueue[T]) Size() int {
	return p.heap.Size()
}

func (p *PriorityQueue[T]) Clear() {
	p.heap.Clear()
}

func (p *PriorityQueue[T]) String() string {
	return p.heap.String()
}
