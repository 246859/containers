package lists

import (
	"fmt"
	"github.com/246859/containers"
	"slices"
)

var _ List[any] = (*ArrayList[any])(nil)

// NewArrayList returns a new ArrayList with the given capacity
func NewArrayList[T any](capacity int, equal containers.Equal[T]) *ArrayList[T] {
	list := &ArrayList[T]{
		elems: make([]T, 0, capacity),
		equal: equal,
	}
	return list
}

type ArrayList[T any] struct {
	elems []T
	equal containers.Equal[T]
}

func (a *ArrayList[T]) Get(i int) (_ T, _ bool) {
	if !a.checkI(i) {
		return
	}
	return a.elems[i], true
}

func (a *ArrayList[T]) Set(i int, elem T) {
	if !a.checkI(i) {
		return
	}
	a.elems[i] = elem
}

func (a *ArrayList[T]) IndexOf(elem T) int {
	for i, e := range a.elems {
		if a.equal(e, elem) {
			return i
		}
	}
	return -1
}

func (a *ArrayList[T]) Add(elems ...T) {
	if len(elems) == 0 {
		return
	}
	a.elems = append(a.elems, elems...)
}

func (a *ArrayList[T]) Insert(i int, elem ...T) {
	if !a.checkI(i) || len(elem) == 0 {
		return
	}
	a.elems = slices.Insert(a.elems, i, elem...)
}

func (a *ArrayList[T]) Remove(i int) {
	if a.checkI(i) {
		a.elems = slices.Delete(a.elems, i, i+1)
	}
}

func (a *ArrayList[T]) RemoveElem(elem T) {
	i := a.IndexOf(elem)
	if i > -1 {
		a.Remove(i)
	}
}

func (a *ArrayList[T]) Contains(elem T) bool {
	return a.IndexOf(elem) > -1
}

func (a *ArrayList[T]) Iterator(reverse bool) containers.IndexIterator[T] {
	size := a.Size()
	snapshot := make([]T, size)
	copy(snapshot, a.elems[:size])
	return containers.NewSliceIndexIterator(reverse, snapshot...)
}

func (a *ArrayList[T]) Values() (_ []T) {
	if a.Size() == 0 {
		return
	}
	var vs []T
	it := a.Iterator(false)
	for it.Next() {
		vs = append(vs, it.Value())
	}
	return vs
}

func (a *ArrayList[T]) Size() int {
	return len(a.elems)
}

func (a *ArrayList[T]) Clear() {
	a.elems = a.elems[:0:0]
}

func (a *ArrayList[T]) Clone() List[T] {
	size := a.Size()
	snapshot := make([]T, size)
	copy(snapshot, a.elems[:size])
	return &ArrayList[T]{
		equal: a.equal,
		elems: snapshot,
	}
}

func (a *ArrayList[T]) Join(list List[T]) {
	if list.Size() == 0 {
		return
	}

	it := list.Iterator(false)
	for it.Next() {
		a.Add(it.Value())
	}
}

func (a *ArrayList[T]) String() string {
	return fmt.Sprintf("%+v", a.elems)
}

func (a *ArrayList[T]) checkI(i int) bool {
	return i >= 0 && i < len(a.elems)
}
