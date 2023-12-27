package stacks

import (
	"fmt"
	"github.com/246859/containers"
	"github.com/246859/containers/lists"
	"strings"
)

var _ Stack[any] = (*LinkedStack[any])(nil)

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		list: lists.NewLinkedList[T](),
	}
}

// LinkedStack is the stack implemented by lists.LinkedList
type LinkedStack[T any] struct {
	list *lists.LinkedList[T]
}

func (l *LinkedStack[T]) Push(t ...T) {
	l.list.Add(t...)
}

func (l *LinkedStack[T]) Pop() (T, bool) {
	val, found := l.list.Get(l.Size() - 1)
	if !found {
		return val, false
	}
	l.list.Remove(l.Size() - 1)
	return val, true
}

func (l *LinkedStack[T]) Peek() (T, bool) {
	return l.list.Get(l.Size() - 1)
}

func (l *LinkedStack[T]) Values() []T {
	return l.list.Values()
}

func (l *LinkedStack[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedStack[T]) Clear() {
	l.list.Clear()
}

func (l *LinkedStack[T]) Iterator() containers.IndexIterator[T] {
	return l.list.Iterator()
}

func (l *LinkedStack[T]) String() string {
	var b strings.Builder
	b.WriteString("LinkedStack[")
	elems := l.Values()
	for i, v := range elems {
		b.WriteString(fmt.Sprintf("%v", v))
		if i != len(elems)-1 {
			b.WriteByte(',')
		}
	}
	b.WriteByte(']')
	return b.String()
}
