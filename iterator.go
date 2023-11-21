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
