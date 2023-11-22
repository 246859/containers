package containers

// Iterable is the base interface of all data structures that can be iterated over
type Iterable[K any, V any] interface {
	Iterator() Iterator[K, V]
}

// IndexIterable is the base interface of all data structures that can be iterated over by slice index
type IndexIterable[V any] interface {
	Iterator() IndexIterator[V]
}

// Iterator is the base interface of the all data structures iterators
// iterator is readonly, represents of data snapshot at a certain moment
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
type IndexIterator[V any] interface {
	Iterator[int, V]
}
