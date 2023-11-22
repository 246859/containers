package containers

// Iterator is the base interface of the all data structures iterators
// K is the index type
// V is the value type
type Iterator[K any, V any] interface {
	// Rewind set the iterator index to the initial state
	Rewind()

	// Reverse changes the iterator orientation
	Reverse()

	// Next returns iterator has reached the end of the container
	Next() bool

	// Index returns iterator current index
	Index() K

	// Value returns value of iterator current index
	Value() V

	// SeekTo move current index to the given index, returns true if success, or returns false
	SeekTo(index K) bool
}

// IndexIterator is base interface of iterator which use slice index
type IndexIterator[T any] interface {
	Iterator[int, T]
}

func NewArrayIndexIterator[T any](reverse bool, elems ...T) *ArrayIndexIterator[T] {
	it := &ArrayIndexIterator[T]{
		reverse: reverse,
		elems:   elems,
	}
	it.Rewind()
	return it
}

var _ IndexIterator[any] = (*ArrayIndexIterator[any])(nil)

type ArrayIndexIterator[T any] struct {
	elems   []T
	index   int
	reverse bool
}

func (a *ArrayIndexIterator[T]) Rewind() {
	a.index = -1
	if a.reverse {
		a.index = len(a.elems)
	}
}

func (a *ArrayIndexIterator[T]) Reverse() {
	a.reverse = !a.reverse
}

func (a *ArrayIndexIterator[T]) Next() bool {
	if !a.reverse && a.index < len(a.elems) {
		a.index++
		return a.index < len(a.elems)
	} else if a.reverse && a.index >= 0 {
		a.index--
		return a.index >= 0
	}
	return false
}

func (a *ArrayIndexIterator[T]) Index() int {
	return a.index
}

func (a *ArrayIndexIterator[T]) Value() T {
	return a.elems[a.Index()]
}

func (a *ArrayIndexIterator[T]) SeekTo(index int) bool {
	if index >= 0 && index < len(a.elems) {
		a.index = index
		return true
	}
	return false
}
