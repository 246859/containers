package lists

import (
	"fmt"
	"github.com/246859/containers"
)

var _ List[any] = (*SinglyLinkedList[any])(nil)

func NewSinglyLinkedList[T any](equal containers.Equal[T]) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{
		equal: equal,
	}
	return list
}

type SinglyLinkedList[T any] struct {
	first *Element[T]
	last  *Element[T]
	equal containers.Equal[T]
	size  int
}

// Element struct of element
type Element[T any] struct {
	value T
	next  *Element[T]
}

func (s *SinglyLinkedList[T]) Get(i int) (_ T, _ bool) {
	if !s.withinRange(i) {
		return
	}
	element := s.first
	for e := 0; e != i; e, element = e+1, element.next {
	}
	return element.value, true
}

func (s *SinglyLinkedList[T]) IndexOf(elem T) int {
	if s.size == 0 {
		return -1
	}
	for index, element := range s.Values() {
		if s.equal(element, elem) {
			return index
		}
	}
	return -1
}

func (s *SinglyLinkedList[T]) Set(i int, elem T) {
	if !s.withinRange(i) {
		// Append
		if i == s.size {
			s.Add(elem)
		}
		return
	}
	element := s.first
	for e := 0; e != i; e, element = e+1, element.next {
	}
	element.value = elem
}

func (s *SinglyLinkedList[T]) Add(elem ...T) {
	for _, e := range elem {
		newElement := &Element[T]{value: e, next: nil}
		if s.size == 0 {
			s.first = newElement
		} else {
			s.last.next = newElement
		}
		s.last = newElement
		s.size++
	}
}

func (s *SinglyLinkedList[T]) Insert(i int, elem ...T) {
	if !s.withinRange(i) {
		// Append
		if i == s.size {
			s.Add(elem...)
		}
		return
	}

	s.size += len(elem)

	var beforeElement *Element[T]
	foundElement := s.first
	for e := 0; e != i; e, foundElement = e+1, foundElement.next {
		beforeElement = foundElement
	}

	if foundElement == s.first {
		oldNextElement := s.first
		for i, value := range elem {
			newElement := &Element[T]{value: value}
			if i == 0 {
				s.first = newElement
			} else {
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range elem {
			newElement := &Element[T]{value: value}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	}

}

func (s *SinglyLinkedList[T]) Remove(i int) {
	if !s.withinRange(i) {
		return
	}

	if s.size == 1 {
		s.Clear()
		return
	}

	var beforeElement *Element[T]
	element := s.first
	for e := 0; e != i; e, element = e+1, element.next {
		beforeElement = element
	}

	if element == s.first {
		s.first = element.next
	}
	if element == s.last {
		s.last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.next = element.next
	}

	element = nil
	s.size--
}

func (s *SinglyLinkedList[T]) RemoveElem(elem T) {
	for i, v := range s.Values() {
		if s.equal(v, elem) {
			s.Remove(i)
		}
	}
}

func (s *SinglyLinkedList[T]) Contains(elem T) bool {
	return s.IndexOf(elem) > -1
}

func (s *SinglyLinkedList[T]) Clone() List[T] {
	sll := NewSinglyLinkedList[T](s.equal)
	sll.Add(s.Values()...)
	return sll
}

func (s *SinglyLinkedList[T]) Join(list List[T]) {
	if list.Size() == 0 {
		return
	}

	it := list.Iterator()
	for it.Next() {
		s.Add(it.Value())
	}
}

func (s *SinglyLinkedList[T]) Iterator() containers.IndexIterator[T] {
	return newListIterator[T](s)
}

func (s *SinglyLinkedList[T]) Values() (_ []T) {
	vs := make([]T, s.size)
	for e, element := 0, s.first; element != nil; e, element = e+1, element.next {
		vs[e] = element.value
	}
	return vs

}

func (s *SinglyLinkedList[T]) Size() int {
	return s.size
}

func (s *SinglyLinkedList[T]) Clear() {
	s.first = nil
	s.last = nil
	s.size = 0
}

func (s *SinglyLinkedList[T]) String() string {
	return fmt.Sprintf("%+v", s.Values())
}

// Check that the index is within bounds of the list
func (s *SinglyLinkedList[T]) withinRange(index int) bool {
	return index >= 0 && index < s.size
}
