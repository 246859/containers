package containers

import "cmp"

const (
	LessThan = -1 + iota
	EqualTo
	GreaterThan
)

// Compare compares e with self, if a is less than b, returns LessThan
// if a is greater than b, returns GreaterThan
// if a is equal to b, returns EqualTo
type Compare[T any] func(a, b T) int

func OrderedCompare[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

// Less check if a is less than b
type Less[T any] func(a, b T) bool

func OrderedLess[T cmp.Ordered](a, b T) bool {
	return cmp.Less(a, b)
}

// Equal checks if a is equals to b
type Equal[T any] func(a, b T) bool

func OrderedEqual[T cmp.Ordered](a, b T) bool {
	return cmp.Compare(a, b) == EqualTo
}
