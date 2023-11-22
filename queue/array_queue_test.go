package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayQueue_Push_Pop(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	queue := NewArrayQueue[int](32)
	queue.Push(data[:5]...)
	queue.Push(data[5:]...)

	for _, d := range data {
		pop, b := queue.Pop()
		assert.True(t, b)
		assert.Equal(t, d, pop)
	}
}

func TestArrayQueue_Push_Peek(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	queue := NewArrayQueue[int](32)
	queue.Push(data[:5]...)
	queue.Push(data[5:]...)

	for _, d := range data {
		pop, b := queue.Peek()
		assert.True(t, b)
		assert.Equal(t, d, pop)

		queue.Pop()
	}
}

func TestArrayQueue_Iterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	queue := NewArrayQueue[int](32)
	queue.Push(data[:5]...)
	queue.Push(data[5:]...)

	it := queue.Iterator(false)
	for it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}

	it.Reverse()
	it.Rewind()

	for it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}
}

func TestArrayQueue_Values(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	queue := NewArrayQueue[int](32)
	queue.Push(data[:5]...)
	queue.Push(data[5:]...)

	values := queue.Values()

	for i := 0; i < len(values); i++ {
		assert.Equal(t, data[i], values[i])
	}
}

func TestArrayQueue_Size(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	queue := NewArrayQueue[int](32)
	queue.Push(data[:5]...)
	assert.Equal(t, 5, queue.Size())
	queue.Push(data[5:]...)
	assert.Equal(t, 10, queue.Size())
}
