package heaps

import (
	"github.com/246859/containers"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestBinaryHeap_Push(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	heap := NewBinaryHeap[int](32, containers.OrderedCompare[int])
	heap.Push(data...)

	ele, has := heap.Pop()
	assert.True(t, has)
	assert.Equal(t, -1, ele)

	ele1, has1 := heap.Pop()
	assert.True(t, has1)
	assert.Equal(t, 1, ele1)

	heap.Push(-100)
	ele2, b2 := heap.Pop()
	assert.True(t, b2)
	assert.Equal(t, -100, ele2)
}

func TestBinaryHeap_Peek(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	heap := NewBinaryHeap[int](32, containers.OrderedCompare[int])
	heap.Push(data...)

	ele, has := heap.Peek()
	assert.True(t, has)
	assert.Equal(t, -1, ele)

	heap.Pop()

	ele1, has1 := heap.Peek()
	assert.True(t, has1)
	assert.Equal(t, 1, ele1)
}

func TestBinaryHeap_Values(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	heap := NewBinaryHeap[int](32, containers.OrderedCompare[int])
	heap.Push(data...)

	values := heap.Values()
	assert.Equal(t, len(data), len(values))
	assert.Equal(t, slices.Min(data), values[0])
}

func TestBinaryHeap_Iterator(t *testing.T) {
	data := []int{2, 1, 4, 6, -1}
	heap := NewBinaryHeap[int](32, containers.OrderedCompare[int])
	heap.Push(data...)

	it := heap.Iterator(false)
	for it.Next() {
		it.Value()
	}

	it.Reverse()
	it.Rewind()

	for it.Next() {
		it.Value()
	}
}
