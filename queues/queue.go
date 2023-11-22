package queues

import "github.com/246859/containers"

// Queue is the base interface of all queues implementations
type Queue[T any] interface {
	// Push pushes an element into queue
	Push(e ...T)
	// Peek returns an element on the head of queue
	Peek() (T, bool)
	// Pop returns an element on the head of queue, then pop it from queue
	Pop() (T, bool)

	containers.IndexIterable[T]

	containers.Container[T]
}
