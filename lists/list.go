package lists

import (
	"github.com/246859/containers"
)

// List is the base interface of all lists
type List[T any] interface {
	// Get returns the elem at the given index
	Get(i int) (T, bool)
	// IndexOf returns the index of first element that matches the given element. if not found, return -1
	IndexOf(elem T, equal containers.Equal[T]) int
	// Set replaces the elem at the given index
	Set(i int, elem T)
	// Add appends the given element to the end of the list
	Add(elem ...T)
	// Insert inserts the given elems into the list at the given index
	Insert(i int, elem ...T)
	// Remove removes the elem match given index from the list
	Remove(i int)
	// RemoveElem removes the first element matches the given from the list,
	RemoveElem(elem T, equal containers.Equal[T])
	// Contains check if list contains the given elements
	Contains(elem T, equal containers.Equal[T]) bool
	// Clone returns a value-copy of the original list, not the deep-copy list
	Clone() List[T]
	// Join joins the given list into the original list
	Join(list List[T])

	containers.IndexIterable[T]

	containers.Container[T]
}

// Swap swaps the two given elements of index in the give list
func Swap[T any](list List[T], i, j int) {
	if i < 0 || j < 0 || i >= list.Size() || j >= list.Size() {
		return
	}

	iv, ei := list.Get(i)
	jv, ej := list.Get(j)

	if !ej || !ei {
		return
	}

	list.Set(i, jv)
	list.Set(j, iv)
}
