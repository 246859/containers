package heaps

import "github.com/246859/containers"

// Heap is the base interface of all heap implementations
type Heap[T any] interface {
	// Push pushes an element into queue
	Push(e ...T)
	// Peek returns an element on the head of queue
	Peek() (T, bool)
	// Pop returns an element on the head of queue, then pop it from queue
	Pop() (T, bool)
	// Fix replace the element at the given index to k, returns if k equal to the element or re-heapify
	Fix(i int, k T)
	// Remove removes an element at the give index from heap
	Remove(i int)
	// Merge merges another heap into this one
	Merge(heap Heap[T])

	containers.IndexIterable[T]

	containers.Container[T]
}
