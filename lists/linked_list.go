package lists

import (
	"fmt"
	"github.com/246859/containers"
	"strings"
)

var _ List[any] = (*LinkedList[any])(nil)

// linked list node
type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// LinkedList implements the list interface by linked nodes which is more efficient to insert and remove nodes.
type LinkedList[T any] struct {
	size int

	first *node[T]
	last  *node[T]
}

// iterate all nodes in the list in give direction, and apply the handler on each node
func (l *LinkedList[T]) iterate(reverse bool, handler func(index int, n *node[T]) bool) {

	if l.size == 0 {
		return
	}

	ni := l.first
	index := 0

	if reverse {
		index = l.size - 1
		ni = l.last
	}

	for ni != nil {

		// apply handler
		if !handler(index, ni) {
			return
		}

		// move to the next
		if !reverse {
			ni = ni.next
			index++
		} else {
			ni = ni.prev
			index--
		}
	}

	return
}

func (l *LinkedList[T]) Get(i int) (val T, found bool) {
	if l.size == 0 || i < 0 || i >= l.size {
		return
	}

	// decide which direction of the iteration cost less move times
	reverse := i+1 > l.size-i

	l.iterate(reverse, func(index int, n *node[T]) bool {
		if i == index {
			val = n.value
			found = true

			return false
		}
		return true
	})

	return
}

func (l *LinkedList[T]) IndexOf(elem T, equal containers.Equal[T]) int {

	foundIndex := -1
	// always iterate in forward direction
	l.iterate(true, func(index int, n *node[T]) bool {
		if equal(elem, n.value) {
			foundIndex = index
			return false
		}
		return true
	})

	return foundIndex
}

func (l *LinkedList[T]) Set(i int, elem T) {

	if l.size == 0 || i < 0 || i >= l.size {
		return
	}

	// decide which direction of the iteration cost less move times
	reverse := i+1 > l.size-i

	l.iterate(reverse, func(index int, n *node[T]) bool {
		if i == index {
			n.value = elem
			return false
		}
		return true
	})

	return
}

func (l *LinkedList[T]) Add(elems ...T) {
	l.Insert(l.size, elems...)
}

func (l *LinkedList[T]) Insert(i int, elems ...T) {

	if i < 0 || (l.size > 0 && i > l.size) {
		return
	}

	var ni *node[T]
	if i == 0 {
		ni = &node[T]{next: l.first}
	} else if i == l.size {
		ni = l.last
	} else {
		reverse := i+1 > l.size-i
		// find the index
		l.iterate(reverse, func(index int, n *node[T]) bool {
			if index == i {
				ni = n.prev
				return false
			}
			return true
		})
	}

	if ni == nil {
		return
	}

	last := ni

	// link the elements
	for _, elem := range elems {
		n := &node[T]{value: elem}

		n.prev = last
		n.next = last.next
		if last.next != nil {
			last.next.prev = n
		}
		last.next = n

		last = n
	}

	// update state
	if i == 0 {
		l.first = ni.next
		l.first.prev = nil
		if l.size == 0 {
			l.last = last
		}
	} else if i == l.size {
		l.last = last
	}

	l.size += len(elems)
}

func (l *LinkedList[T]) unlink(n *node[T]) {
	if n == nil {
		return
	}

	next := n.next
	prev := n.prev

	if prev != n {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	n.prev = nil
	n.next = nil
}

func (l *LinkedList[T]) Remove(i int) {

	if l.size == 0 || i < 0 || i >= l.size {
		return
	}

	var ni *node[T]
	reverse := i+1 > l.size-i
	// find the index
	l.iterate(reverse, func(index int, n *node[T]) bool {
		if index == i {
			ni = n
			return false
		}
		return true
	})

	// unlink the node
	next := ni.next
	prev := ni.prev

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	if i == 0 {
		l.first = next
	} else if i == l.size-1 {
		l.last = prev
	}

	ni.prev = nil
	ni.next = nil

	l.size--
}

func (l *LinkedList[T]) RemoveElem(elem T, equal containers.Equal[T]) {
	if l.size == 0 {
		return
	}

	var ni *node[T]
	i := -1
	// find the element
	l.iterate(true, func(index int, n *node[T]) bool {
		if equal(elem, n.value) {
			ni = n
			i = index
			return false
		}
		return true
	})

	// unlink the node
	next := ni.next
	prev := ni.prev

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	if i == 0 {
		l.first = next
	} else if i == l.size-1 {
		l.last = prev
	}

	ni.prev = nil
	ni.next = nil
	l.size--
}

func (l *LinkedList[T]) Contains(elem T, equal containers.Equal[T]) bool {
	return l.IndexOf(elem, equal) != -1
}

func (l *LinkedList[T]) Clone() List[T] {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[T]) Join(list List[T]) {
	if list.Size() == 0 {
		return
	}

	var elems []T
	for it := list.Iterator(); it.Valid(); it.Next() {
		elems = append(elems, it.Value())
	}
	l.Add(elems...)
}

func (l *LinkedList[T]) Iterator() containers.IndexIterator[T] {
	it := &linkedIterator[T]{
		list: l,
		size: l.size,
	}
	it.Rewind()

	return it
}

func (l *LinkedList[T]) Values() []T {
	var elems []T
	for it := l.Iterator(); it.Valid(); it.Next() {
		elems = append(elems, it.Value())
	}
	return elems
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}

func (l *LinkedList[T]) String() string {
	var b strings.Builder
	b.WriteString("LinkedList[")
	elems := l.Values()
	for i, v := range l.Values() {
		b.WriteString(fmt.Sprintf("%v", v))
		if i != len(elems)-1 {
			b.WriteByte(',')
		}
	}
	b.WriteByte(']')
	return b.String()
}

// LinkedList iterator
type linkedIterator[T any] struct {
	list *LinkedList[T]
	size int

	index int
	// record index of the node
	indexNode *node[T]

	reverse bool
}

func (it *linkedIterator[T]) Rewind() {
	it.index = 0
	it.indexNode = it.list.first

	if it.reverse {
		it.index = it.size - 1
		it.indexNode = it.list.last
	}
}

func (it *linkedIterator[T]) Reverse() {
	it.reverse = !it.reverse
}

func (it *linkedIterator[T]) Valid() bool {
	return it.index >= 0 && it.index < it.size && it.indexNode != nil
}

func (it *linkedIterator[T]) Next() {
	if !it.Valid() {
		return
	}

	step := 1
	next := it.indexNode.next

	if it.reverse {
		step = -1
		next = it.indexNode.prev
	}

	it.index += step
	it.indexNode = next
}

func (it *linkedIterator[T]) Index() int {
	if !it.Valid() {
		return -1
	}
	return it.index
}

func (it *linkedIterator[T]) Value() T {
	var v T
	if it.Valid() {
		v = it.indexNode.value
	}
	return v
}

func (it *linkedIterator[T]) SeekTo(index int) bool {
	if index < 0 || index >= it.size {
		return false
	}

	if index == it.index {
		return true
	}

	oriR := it.reverse

	// decide the seek direction
	it.reverse = index < it.index
	for it.index != index {
		it.Next()
	}
	it.reverse = oriR

	return true
}
