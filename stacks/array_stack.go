package stacks

import (
	"fmt"
	"github.com/246859/containers"
	"github.com/246859/containers/internal"
)

var _ Stack[any] = (*ArrayStack[any])(nil)

type ArrayStack[T any] struct {
	s []T
}

func NewArrayStack[T any](capacity int) *ArrayStack[T] {
	return &ArrayStack[T]{s: make([]T, 0, capacity)}
}

func (stack *ArrayStack[T]) Push(t ...T) {
	stack.s = append(stack.s, t...)
}

func (stack *ArrayStack[T]) Pop() (_ T, _ bool) {
	if stack.Size() == 0 {
		return
	}
	peek := stack.s[len(stack.s)-1]
	stack.s = stack.s[:len(stack.s)-1]
	return peek, true
}

func (stack *ArrayStack[T]) Peek() (_ T, _ bool) {
	if stack.Size() == 0 {
		return
	}
	return stack.s[len(stack.s)-1], true
}
func (stack *ArrayStack[T]) Iterator() containers.IndexIterator[T] {
	return internal.NewSliceIterator(stack.s[:stack.Size()])
}

func (stack *ArrayStack[T]) Values() []T {
	var vals []T

	for it := stack.Iterator(); it.Valid(); it.Next() {
		vals = append(vals, it.Value())
	}
	return vals
}

func (stack *ArrayStack[T]) Size() int {
	return len(stack.s)
}

func (stack *ArrayStack[T]) Clear() {
	stack.s = stack.s[:0:0]
}

func (stack *ArrayStack[T]) String() string {
	return fmt.Sprintf("%+v", stack.s)
}
