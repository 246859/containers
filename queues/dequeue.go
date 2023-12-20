package queues

import (
	"fmt"
	"github.com/246859/containers"
	"strings"
)

const (
	dequeueGowThreshold = 0.8
)

var _ Queue[any] = (Deque[any])(nil)

// Deque represents the double-end queue, which could push item and pop item in both left and right end.
type Deque[T any] interface {
	Queue[T]
	// LPush push an item to the queue left end
	LPush(es ...T)
	// RPeek returns the first item from the queue right end
	RPeek() (T, bool)
	// RPop pop the first item from the queue right end
	RPop() (T, bool)
}

var _ Deque[any] = (*ArrayDeque[any])(nil)

func NewArrayDequeue[T any](size int) *ArrayDeque[T] {
	if size <= 0 {
		size = 128
	}

	return &ArrayDeque[T]{
		capacity: size,
		arr:      make([]T, size),
		front:    0,
		tail:     0,
	}
}

// ArrayDeque is the array implementation of Deque,
// its ability of dynamically growing depends on the go slice implementation.
type ArrayDeque[T any] struct {
	arr      []T
	capacity int

	size  int
	front int
	tail  int
}

// Push represents RPush, push an item to the queue right end
func (queue *ArrayDeque[T]) Push(es ...T) {
	queue.resize(true)
	for _, e := range es {
		queue.arr[queue.tail] = e
		queue.tail = (queue.tail + 1) % queue.capacity
		queue.size++
	}
}

// Peek represents RPeek, returns the first item of the queue right end
func (queue *ArrayDeque[T]) Peek() (_ T, _ bool) {
	if queue.size == 0 {
		return
	}
	return queue.arr[queue.front], true
}

// Pop represents LPop, pop the first item of the queue left end
func (queue *ArrayDeque[T]) Pop() (_ T, _ bool) {
	peek, ok := queue.Peek()
	if !ok {
		return
	}
	queue.front = (queue.front + 1) % queue.capacity
	queue.size--

	queue.resize(false)
	return peek, true
}

// LPush pushed an item to the queue left end
func (queue *ArrayDeque[T]) LPush(es ...T) {
	queue.resize(true)
	for _, e := range es {
		queue.front = (queue.front - 1 + queue.capacity) % queue.capacity
		queue.arr[queue.front] = e
		queue.size++
	}
}

// RPeek returns the first item from queue right end.
func (queue *ArrayDeque[T]) RPeek() (_ T, _ bool) {
	if queue.size == 0 {
		return
	}
	return queue.arr[queue.tail-1], true
}

// RPop pops the first item from queue right end.
func (queue *ArrayDeque[T]) RPop() (_ T, _ bool) {
	peek, ok := queue.RPeek()
	if !ok {
		return
	}
	queue.tail = (queue.tail - 1 + queue.capacity) % queue.capacity
	queue.size--

	queue.resize(false)
	return peek, true
}

// gow queue capacity to hold more elements
func (queue *ArrayDeque[T]) resize(grow bool) {
	var newCapacity int
	// size > capacity * 0.8, growing
	if grow && float64(queue.size) >= float64(queue.capacity)*dequeueGowThreshold {
		// if capacity < 256, grow queue capacity to double original capacity
		//  otherwise grow to 5/4 original capacity
		var growCap int
		if queue.capacity < 256 {
			growCap = queue.capacity
		} else {
			growCap = queue.capacity >> 2
		}

		newCapacity = queue.capacity + growCap
		// size > 256 && size < capacity/2, shrinking
	} else if !grow && queue.capacity >= 256 && queue.size < queue.capacity>>1 {
		// shrink queue capacity to 3/4 original capacity,
		// a small shrink step could avoid growing frequently.
		shrinkCap := queue.capacity >> 2
		newCapacity = queue.capacity - shrinkCap
	} else {
		return
	}

	newArr := make([]T, newCapacity)

	// copy right end elements to new arr left end
	copy(newArr[:queue.tail], queue.arr[:queue.tail])

	// no need to copy if left end has no elements
	if queue.front > 0 || queue.tail == 0 {
		newFront := newCapacity - (queue.capacity - queue.front)
		// copy the elements of left end of queue to new arr right end
		copy(newArr[newFront:newCapacity], queue.arr[queue.front:queue.capacity])
		queue.front = newFront
	}

	queue.arr = newArr
	queue.capacity = newCapacity
}

func (queue *ArrayDeque[T]) Iterator() containers.IndexIterator[T] {
	return &arrayDeQueIterator[T]{
		queue: queue,
		size:  queue.capacity,
		l:     queue.front,
		r:     queue.tail,
		index: queue.front,
	}
}

func (queue *ArrayDeque[T]) Values() []T {
	var vals []T
	for it := queue.Iterator(); it.Valid(); it.Next() {
		vals = append(vals, it.Value())
	}
	return vals
}

func (queue *ArrayDeque[T]) Size() int {
	return queue.size
}

func (queue *ArrayDeque[T]) Clear() {
	queue.arr = queue.arr[:0:0]
}

func (queue *ArrayDeque[T]) String() string {
	var bf strings.Builder
	bf.WriteString("ArrayDeque[")
	var elems []string
	for it := queue.Iterator(); it.Valid(); it.Next() {
		elems = append(elems, fmt.Sprintf("%v", it.Value()))
	}
	bf.WriteString(strings.Join(elems, ","))
	bf.WriteString("]")
	return bf.String()
}

// arrayDeQueIterator is an iterator of ArrayDeque,
// iterate over the elements in range [queue.front, queue.tail]
type arrayDeQueIterator[T any] struct {
	queue *ArrayDeque[T]

	size int
	// front of queue
	l int
	// tail of queue
	r int

	index   int
	reverse bool
}

func (it *arrayDeQueIterator[T]) Rewind() {
	it.index = it.l
	if it.reverse {
		it.index = it.r - 1
	}
}

func (it *arrayDeQueIterator[T]) Reverse() {
	it.reverse = !it.reverse
}

func (it *arrayDeQueIterator[T]) Valid() bool {
	if it.l > it.r {
		return it.index < it.r || it.index >= it.l
	} else {
		return it.index >= it.l && it.index < it.r
	}
}

func (it *arrayDeQueIterator[T]) Next() {
	if !it.Valid() {
		return
	}

	if it.reverse {
		it.index = (it.index - 1 + it.size) % it.size
	} else {
		it.index = (it.index + 1) % it.size
	}
}

func (it *arrayDeQueIterator[T]) Index() int {
	return it.index
}

func (it *arrayDeQueIterator[T]) Value() (_ T) {
	if !it.Valid() {
		return
	}
	return it.queue.arr[it.index]
}

func (it *arrayDeQueIterator[T]) SeekTo(index int) bool {
	if index >= it.l && index < it.r {
		it.index = index
		return true
	}
	return false
}
