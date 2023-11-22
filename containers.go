package containers

// Container is the base interface represent of all data structures
type Container[T any] interface {
	// Values returns all the values of the container
	Values() []T
	// Size returns the size of the container elements
	Size() int
	// Clear clears all the elements of the container
	Clear()
	// String returns the string representation of the container
	String() string
}
