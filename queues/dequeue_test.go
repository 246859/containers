package queues

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestArrayDeque_New(t *testing.T) {
	dequeue := NewArrayDequeue[int](64)
	assert.NotNil(t, dequeue)
}

func TestArrayDeque_Push(t *testing.T) {
	dequeue := NewArrayDequeue[int](10)
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

func TestArrayDeque_Push_Pop(t *testing.T) {
	dequeue := NewArrayDequeue[int](10)
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

func TestArrayDeque_Push_Pop_Grow(t *testing.T) {
	{
		dequeue := NewArrayDequeue[int](10)
		for i := 0; i < 11; i++ {
			dequeue.Push(i)
		}
		t.Log(dequeue.Values())
		assert.EqualValues(t, 11, dequeue.Size())
		assert.EqualValues(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, dequeue.Values())
	}
	{
		dequeue := NewArrayDequeue[int](10)
		for i := 0; i < 11; i++ {
			dequeue.LPush(i)
		}
		t.Log(dequeue.Values())
		assert.EqualValues(t, 11, dequeue.Size())
		assert.EqualValues(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, dequeue.Values())
	}
	{
		dequeue := NewArrayDequeue[int](10)
		for i := 0; i < 11; i++ {
			if i%2 == 0 {
				dequeue.Push(i)
			} else {
				dequeue.LPush(i)
			}
		}
		t.Log(dequeue.Values())
		assert.EqualValues(t, 11, dequeue.Size())
		assert.EqualValues(t, []int{9, 7, 5, 3, 1, 0, 2, 4, 6, 8, 10}, dequeue.Values())
	}
	{
		dequeue := NewArrayDequeue[int](10)
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
	}
}

func TestArrayDeque_Push_Pop_Shrink(t *testing.T) {
	{
		dequeue := NewArrayDequeue[int](10)
		var expectedCount = 300
		var expected []int
		for i := 0; i < expectedCount; i++ {
			if i < expectedCount/2 {
				expected = append(expected, i)
				dequeue.Push(i)
			} else {
				expected = slices.Insert(expected, 0, i-expectedCount/2)
				dequeue.LPush(i - expectedCount/2)
			}
		}
		t.Log(dequeue.Values())
		assert.EqualValues(t, expectedCount, dequeue.Size())
		assert.EqualValues(t, expected, dequeue.Values())

		for i := 0; i < expectedCount/2; i++ {
			dequeue.Pop()
			expected = expected[1:]
		}
		t.Log(dequeue.capacity)
		t.Log(dequeue.Size())
		t.Log(dequeue.Values())
		assert.EqualValues(t, expected, dequeue.Values())
	}
	{
		dequeue := NewArrayDequeue[int](10)
		var expectedCount = 300
		var expected []int
		for i := 0; i < expectedCount; i++ {
			if i < expectedCount/2 {
				expected = append(expected, i)
				dequeue.Push(i)
			} else {
				expected = slices.Insert(expected, 0, i-expectedCount/2)
				dequeue.LPush(i - expectedCount/2)
			}
		}
		t.Log(dequeue.Values())
		assert.EqualValues(t, expectedCount, dequeue.Size())
		assert.EqualValues(t, expected, dequeue.Values())

		for i := 0; i < expectedCount/2; i++ {
			dequeue.RPop()
			expected = expected[:len(expected)-1]
		}
		t.Log(dequeue.capacity)
		t.Log(dequeue.Size())
		t.Log(dequeue.Values())
		assert.EqualValues(t, expected, dequeue.Values())
	}
}
