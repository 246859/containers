package stacks

import "github.com/246859/containers"

type Stack[T any] interface {
	Push(t ...T)
	Pop() (T, bool)
	Peek() (T, bool)

	containers.Container[T]

	containers.IndexIterable[T]
}
