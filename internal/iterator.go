package internal

import "github.com/246859/containers"

var _ containers.IndexIterator[any] = (*SliceIterator[any])(nil)

func NewSliceIterator[T any](elems []T) *SliceIterator[T] {
	it := &SliceIterator[T]{
		elems: elems,
	}
	it.Rewind()
	return it
}

// SliceIterator iterates a slice
type SliceIterator[T any] struct {
	elems   []T
	index   int
	reverse bool
}

func (a *SliceIterator[T]) Rewind() {
	a.index = -1
	if a.reverse {
		a.index = len(a.elems)
	}
}

func (a *SliceIterator[T]) Reverse() {
	a.reverse = !a.reverse
}

func (a *SliceIterator[T]) Next() bool {
	if !a.reverse && a.index < len(a.elems) {
		a.index++
		return a.index < len(a.elems)
	} else if a.reverse && a.index >= 0 {
		a.index--
		return a.index >= 0
	}
	return false
}

func (a *SliceIterator[T]) Index() int {
	return a.index
}

func (a *SliceIterator[T]) Value() T {
	return a.elems[a.Index()]
}

func (a *SliceIterator[T]) SeekTo(index int) bool {
	if index >= 0 && index < len(a.elems) {
		a.index = index
		return true
	}
	return false
}
