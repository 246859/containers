package queues

import (
	"fmt"
	"github.com/246859/containers"
	"github.com/246859/containers/lists"
	"strings"
)

var _ Deque[any] = (*LinkedQueue[any])(nil)

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		list: lists.NewLinkedList[T](),
	}
}

// LinkedQueue implements the Queue interface and Deque interface.
// It is different from ArrayDeque because it is based on lists.LinkedList,
// so no need to consider capacity of growing and shrinking.
type LinkedQueue[T any] struct {
	list *lists.LinkedList[T]
}

func (l *LinkedQueue[T]) Push(es ...T) {
	l.list.Add(es...)
}

func (l *LinkedQueue[T]) Peek() (T, bool) {
	return l.list.Get(0)
}

func (l *LinkedQueue[T]) Pop() (T, bool) {
	val, found := l.list.Get(0)
	if !found {
		return val, false
	}
	l.list.Remove(0)
	return val, true
}

func (l *LinkedQueue[T]) LPush(es ...T) {
	l.list.Insert(0, es...)
}

func (l *LinkedQueue[T]) RPeek() (T, bool) {
	return l.list.Get(l.Size() - 1)
}

func (l *LinkedQueue[T]) RPop() (T, bool) {
	val, found := l.list.Get(l.Size() - 1)
	if !found {
		return val, false
	}
	l.list.Remove(l.Size() - 1)
	return val, true
}

func (l *LinkedQueue[T]) Iterator() containers.IndexIterator[T] {
	return l.list.Iterator()
}

func (l *LinkedQueue[T]) Values() []T {
	return l.list.Values()
}

func (l *LinkedQueue[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedQueue[T]) Clear() {
	l.list.Clear()
}

func (l *LinkedQueue[T]) String() string {
	var b strings.Builder
	b.WriteString("LinkedQueue[")
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
