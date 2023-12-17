package lists

import "github.com/246859/containers"

var _ containers.IndexIterator[any] = (*Iterator[any])(nil)

func newIterator[T any](list List[T]) *Iterator[T] {
	it := &Iterator[T]{
		list: list,
		// record the current size
		size: list.Size(),
	}
	it.Rewind()
	return it
}

// Iterator returns a basic list iterator
type Iterator[T any] struct {
	list    List[T]
	size    int
	index   int
	reverse bool
}

func (i *Iterator[T]) Valid() bool {
	return 0 <= i.index && i.index < i.size
}
func (i *Iterator[T]) Rewind() {
	i.index = 0
	if i.reverse {
		i.index = i.size - 1
	}
}

func (i *Iterator[T]) Reverse() {
	i.reverse = !i.reverse
}

func (i *Iterator[T]) Next() {
	if !i.Valid() {
		return
	}

	if i.reverse {
		i.index--
	} else {
		i.index++
	}
}

func (i *Iterator[T]) Index() int {
	return i.index
}

func (i *Iterator[T]) Value() (_ T) {
	if !i.Valid() {
		return
	}
	v, _ := i.list.Get(i.index)
	return v
}

func (i *Iterator[T]) SeekTo(index int) bool {
	if index < 0 || index >= i.size {
		return false
	}
	i.index = index
	return true
}
