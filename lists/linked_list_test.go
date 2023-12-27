package lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	_ = NewLinkedList[any]()
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4)
	list.Add(5)
	list.Add(6)
	list.Add(7)

	assert.EqualValues(t, 7, list.Size())
}

func TestLinkedList_Add_1(t *testing.T) {
	list := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		list.Add(i)
	}
	assert.EqualValues(t, 1000, list.Size())
	for i := 0; i < 1000; i++ {
		val, found := list.Get(i)
		assert.True(t, found)
		assert.EqualValues(t, i, val)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4, 5)
	list.Insert(0, 6)
	val, found := list.Get(0)
	assert.True(t, found)
	assert.EqualValues(t, 6, val)

	list.Insert(list.size, 9)
	get, f := list.Get(list.size - 1)
	assert.True(t, f)
	assert.EqualValues(t, 9, get)

	list.Insert(1, 7)
	fmt.Println(list)
	v, b := list.Get(1)
	assert.True(t, b)
	assert.EqualValues(t, 7, v)
}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4, 5, 6)

	// remove head
	list.Remove(0)
	assert.EqualValues(t, 5, list.Size())

	val, found := list.Get(0)
	assert.True(t, found)
	assert.EqualValues(t, val, 2)

	// remove tail
	list.Remove(list.Size() - 1)
	get, f := list.Get(list.Size() - 1)

	assert.True(t, f)
	assert.EqualValues(t, 5, get)

	// remove middle
	list.Remove(1)
	v, b := list.Get(1)
	assert.True(t, b)
	assert.EqualValues(t, 4, v)
}

func TestLinkedList_RemoveElem(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4, 5, 6)

	list.RemoveElem(2, func(a, b int) bool {
		return a == b
	})

	val, found := list.Get(1)
	assert.True(t, found)
	assert.EqualValues(t, 3, val)
	assert.EqualValues(t, 5, list.Size())
}

func TestLinkedList_Clear(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4, 5, 6)

	list.Clear()
	assert.EqualValues(t, 0, list.Size())

	list.Add(7)
	val, found := list.Get(0)
	assert.True(t, found)
	assert.EqualValues(t, 7, val)
}

func TestLinkedList_Iterator(t *testing.T) {
	list := NewLinkedList[int]()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}

	i := 0
	for it := list.Iterator(); it.Valid(); it.Next() {
		assert.EqualValues(t, i, it.Value())
		i++
	}
}

func TestLinkedList_Contains(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1, 2, 3, 4, 5, 6)

	contains := list.Contains(1, func(a, b int) bool {
		return a == b
	})

	assert.True(t, contains)

	contains = list.Contains(100, func(a, b int) bool {
		return a == b
	})

	assert.False(t, contains)
}
