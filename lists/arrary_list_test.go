package lists

import (
	"github.com/246859/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitial(t *testing.T) {
	_ = NewArrayList[int](32)
	_ = NewArrayList[string](32)
	type p struct {
		age int
	}
	_ = NewArrayList[p](32)
}

func TestArrayList_Add(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	for i, num := range data {
		e, has := list.Get(i)
		assert.True(t, has)
		assert.Equal(t, num, e)
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	for i, num := range data {
		index := list.IndexOf(num, containers.OrderedEqual[int])
		assert.Equal(t, i, index)
	}

	assert.Equal(t, -1, list.IndexOf(100, containers.OrderedEqual[int]))
}

func TestArrayList_Contains(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	for _, num := range data {
		assert.True(t, list.Contains(num, containers.OrderedEqual[int]))
	}

	assert.False(t, list.Contains(100, containers.OrderedEqual[int]))
}

func TestArrayList_Size(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	assert.Equal(t, len(data), list.Size())

	list.Add(200)
	assert.Equal(t, len(data)+1, list.Size())
}

func TestArrayList_Set(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	{
		list.Set(0, 100)
		get, b := list.Get(0)
		assert.True(t, b)
		assert.Equal(t, 100, get)
	}

	{
		list.Set(1, 202)
		get, b := list.Get(1)
		assert.True(t, b)
		assert.Equal(t, 202, get)
	}
}

func TestArrayList_Insert(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	{
		index, values := 0, []int{-1, -2, -3}
		list.Insert(index, values...)

		for i, value := range values {
			get, b := list.Get(index + i)
			assert.True(t, b)
			assert.Equal(t, value, get)
		}
	}
}

func TestArrayList_Remove(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	list.Remove(0)
	get, b := list.Get(0)
	assert.True(t, b)
	assert.NotEqual(t, 1, get)

	list.RemoveElem(0, containers.OrderedEqual[int])
	list.RemoveElem(2, containers.OrderedEqual[int])
	list.RemoveElem(3, containers.OrderedEqual[int])

	assert.False(t, list.Contains(0, containers.OrderedEqual[int]))
	assert.False(t, list.Contains(2, containers.OrderedEqual[int]))
	assert.False(t, list.Contains(3, containers.OrderedEqual[int]))
}

func TestArrayList_Clone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	coneList := list.Clone()
	assert.Equal(t, list.Size(), coneList.Size())
	assert.Equal(t, list.String(), coneList.String())
}

func TestArrayList_Join(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	data1 := []int{10, 11, 12, 13, 14, 15}
	list1 := NewArrayList[int](32)
	list1.Add(data1...)

	list.Join(list1)

	for i := len(data); i < len(data)+len(data1); i++ {
		get, b := list.Get(i)
		assert.True(t, b)
		assert.Equal(t, data1[i-len(data)], get)
	}
}

func TestArrayList_Clear(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	assert.Equal(t, len(data), list.Size())
	list.Clear()
	assert.Equal(t, 0, list.Size())
}

func TestArrayList_Values(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	values := list.Values()
	assert.NotNil(t, values)
	for i, value := range values {
		assert.Equal(t, data[i], value)
	}
}

func TestArrayList_Iterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := NewArrayList[int](32)
	list.Add(data...)

	it := list.Iterator()
	for ; it.Valid(); it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}

	it.Rewind()
	assert.Equal(t, 0, it.Index())

	it.Reverse()

	it.Rewind()
	assert.Equal(t, len(data)-1, it.Index())
	for ; it.Valid(); it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}

	it.SeekTo(2)
	assert.Equal(t, 2, it.Index())
	assert.Equal(t, data[2], it.Value())
}
