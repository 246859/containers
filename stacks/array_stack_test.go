package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack_Push_Pop(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stack := NewArrayStack[int](32)
	stack.Push(data[:5]...)
	stack.Push(data[5:]...)

	for i := len(data) - 1; i >= 0; i-- {
		d := data[i]
		pop, b := stack.Pop()
		assert.True(t, b)
		assert.Equal(t, d, pop)
	}
}

func TestArrayStack_Push_Peek(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stack := NewArrayStack[int](32)
	stack.Push(data[:5]...)
	stack.Push(data[5:]...)

	for i := len(data) - 1; i >= 0; i-- {
		d := data[i]
		pop, b := stack.Peek()
		assert.True(t, b)
		assert.Equal(t, d, pop)

		stack.Pop()
	}

}

func TestArrayStack_Iterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stack := NewArrayStack[int](32)
	stack.Push(data[:5]...)
	stack.Push(data[5:]...)

	it := stack.Iterator()
	for ; it.Valid(); it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}

	it.Reverse()
	it.Rewind()

	for ; it.Valid(); it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}
}

func TestArrayStack_Values(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stack := NewArrayStack[int](32)
	stack.Push(data[:5]...)
	stack.Push(data[5:]...)

	values := stack.Values()

	for i := 0; i < len(values); i++ {
		assert.Equal(t, data[i], values[i])
	}
}

func TestArrayStack_Size(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stack := NewArrayStack[int](32)
	stack.Push(data[:5]...)
	assert.Equal(t, 5, stack.Size())
	stack.Push(data[5:]...)
	assert.Equal(t, 10, stack.Size())
}
