package queues

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestLinkedQueue_New(t *testing.T) {
	dequeue := NewLinkedQueue[int]()
	assert.NotNil(t, dequeue)
}

func TestLinkedQueue_Peek(t *testing.T) {
	dequeue := NewLinkedQueue[int]()
	// push
	dequeue.Push(1)
	dequeue.LPush(4)
	dequeue.Push(2)
	dequeue.LPush(5)

	lPeek, b := dequeue.Peek()
	assert.True(t, b)
	assert.EqualValues(t, 5, lPeek)

	rPeek, b2 := dequeue.RPeek()
	assert.True(t, b2)
	assert.EqualValues(t, 2, rPeek)
}

func TestLinkedQueue_Push(t *testing.T) {
	dequeue := NewLinkedQueue[int]()
	// push
	dequeue.Push(1)
	dequeue.LPush(4)
	dequeue.Push(2)
	dequeue.LPush(5)

	t.Log(dequeue.String())

	values := dequeue.Values()
	expected := []int{5, 4, 1, 2}
	assert.EqualValues(t, expected, values)
}

func TestLinkedQueue_Push_Pop(t *testing.T) {
	dequeue := NewLinkedQueue[int]()
	// push
	dequeue.Push(1)
	dequeue.LPush(4)
	dequeue.Push(2)
	dequeue.LPush(5)

	t.Log(dequeue.String())

	pop1, ok1 := dequeue.Pop()
	assert.True(t, ok1)
	assert.EqualValues(t, 5, pop1)

	pop2, ok2 := dequeue.RPop()
	assert.True(t, ok2)
	assert.EqualValues(t, 2, pop2)
	t.Log(dequeue.Values())
}

func TestNewLinkedQueue_Iterator(t *testing.T) {
	dequeue := NewLinkedQueue[int]()
	var expectedCount = 1000
	var expected []int
	for i := 0; i < expectedCount; i++ {
		if i%2 == 0 {
			expected = append(expected, i)
			dequeue.Push(i)
		} else {
			expected = slices.Insert(expected, 0, i)
			dequeue.LPush(i)
		}
	}
	assert.EqualValues(t, expectedCount, dequeue.Size())
	assert.EqualValues(t, expected, dequeue.Values())

	for it := dequeue.Iterator(); it.Valid(); it.Next() {
		assert.EqualValues(t, expected[it.Index()], it.Value())
	}
}
