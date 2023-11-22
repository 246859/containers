package queues

import (
	"github.com/246859/containers"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestPriorityQueue_Push_Pop(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	pQueue := NewPriorityQueue[int](32, containers.OrderedCompare[int])
	pQueue.Push(data...)

	ele, has := pQueue.Pop()
	assert.True(t, has)
	assert.Equal(t, -1, ele)

	ele1, has1 := pQueue.Pop()
	assert.True(t, has1)
	assert.Equal(t, 1, ele1)

	pQueue.Push(-100)
	ele2, b2 := pQueue.Pop()
	assert.True(t, b2)
	assert.Equal(t, -100, ele2)
}

func TestPriorityQueue_Peek(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	pQueue := NewPriorityQueue[int](32, containers.OrderedCompare[int])
	pQueue.Push(data...)

	ele, has := pQueue.Peek()
	assert.True(t, has)
	assert.Equal(t, -1, ele)

	pQueue.Pop()

	ele1, has1 := pQueue.Peek()
	assert.True(t, has1)
	assert.Equal(t, 1, ele1)
}

func TestPriorityQueue_Values(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	pQueue := NewPriorityQueue[int](32, containers.OrderedCompare[int])
	pQueue.Push(data...)

	values := pQueue.Values()
	assert.Equal(t, len(data), len(values))
	assert.Equal(t, slices.Min(data), values[0])
}

func TestPriorityQueue_Iterator(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	pQueue := NewPriorityQueue[int](32, containers.OrderedCompare[int])
	pQueue.Push(data...)

	it := pQueue.Iterator(false)
	for it.Next() {
		it.Value()
	}

	it.Reverse()
	it.Rewind()

	for it.Next() {
		it.Value()
	}
}
