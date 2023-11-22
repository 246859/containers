package heaps

import (
	"github.com/246859/containers"
	"github.com/246859/containers/lists"
)

var _ Heap[any] = (*BinaryHeap[any])(nil)

func NewBinaryHeap[T any](capacity int, compare containers.Compare[T]) *BinaryHeap[T] {
	list := lists.NewArrayList[T](capacity, func(a, b T) bool {
		return compare(a, b) == containers.EqualTo
	})

	return &BinaryHeap[T]{
		list: list,
		cmp:  compare,
	}
}

type BinaryHeap[T any] struct {
	list *lists.ArrayList[T]
	cmp  containers.Compare[T]
}

func (heap *BinaryHeap[T]) Push(es ...T) {
	if len(es) == 1 {
		heap.list.Add(es[0])
		heap.up(heap.Size() - 1)
	} else {
		// push one then up one that is the normal method which will run in O(nlogn) time
		// but another faster method as follows that reference https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap
		heap.list.Add(es...)
		// get the last possible subtree root node position
		size := heap.list.Size() / 2
		// iterate over all subtree root node bottom up, and execute down operation in per root node
		// Assuming that the subtrees of height h have all been binary heapified, then for the subtrees of height h+1,
		// adjusting the root node along the branch of the maximum child node requires at most h steps to complete the binary heapification.
		// It can be proven that the time complexity of this algorithm is O(n).
		for i := size; i >= 0; i-- {
			heap.down(i)
		}
	}
}

func (heap *BinaryHeap[T]) Peek() (_ T, _ bool) {
	elem, has := heap.list.Get(0)
	if !has {
		return
	}
	return elem, true
}

func (heap *BinaryHeap[T]) Pop() (_ T, _ bool) {
	elem, has := heap.list.Get(0)
	if !has {
		return
	}
	last := heap.list.Size() - 1
	lists.Swap[T](heap.list, 0, last)
	heap.list.Remove(last)
	heap.down(0)
	return elem, true
}

func (heap *BinaryHeap[T]) Remove(i int) {
	size := heap.Size()
	if i < 0 || i > size {
		return
	}
	n := size - 1
	lists.Swap[T](heap.list, i, n)
	heap.list.Remove(n)
	heap.down(i)
}

func (heap *BinaryHeap[T]) Fix(i int, k T) {
	ele, has := heap.list.Get(i)
	if !has {
		return
	}
	n := heap.Size() - 1

	switch heap.cmp(k, ele) {
	case containers.LessThan:
		heap.list.Set(i, k)
		lists.Swap[T](heap.list, 0, i)
		heap.down(0)
	case containers.GreaterThan:
		heap.list.Set(i, k)
		lists.Swap[T](heap.list, n, i)
		heap.up(n)
	}
}

func (heap *BinaryHeap[T]) Merge(h Heap[T]) {
	heap.Push(h.Values()...)
}

func (heap *BinaryHeap[T]) Iterator(reverse bool) containers.IndexIterator[T] {
	size := heap.Size()
	snapshot := make([]T, size)
	copy(snapshot, heap.list.Values()[:size])
	return containers.NewSliceIndexIterator(reverse, snapshot...)
}

func (heap *BinaryHeap[T]) Values() []T {
	var vals []T
	it := heap.Iterator(false)
	for it.Next() {
		vals = append(vals, it.Value())
	}
	return vals
}

func (heap *BinaryHeap[T]) Size() int {
	return heap.list.Size()
}

func (heap *BinaryHeap[T]) Clear() {
	heap.list.Clear()
}

func (heap *BinaryHeap[T]) String() string {
	return heap.list.String()
}

// up the last element in the heap until it finds a parent node less than itself
func (heap *BinaryHeap[T]) up(i int) {
	if i < 0 || i >= heap.list.Size() {
		return
	}

	// parent = index / 2 - 1
	for pi := (i - 1) >> 1; i > 0; pi = (i - 1) >> 1 {
		v, _ := heap.list.Get(i)
		pv, _ := heap.list.Get(pi)

		if heap.cmp(v, pv) >= containers.EqualTo {
			break
		}

		lists.Swap[T](heap.list, i, pi)
		i = pi
	}
}

// down the element at the top of the heap until it finds a child node greater than itself.
func (heap *BinaryHeap[T]) down(i int) {
	if i < 0 || i >= heap.list.Size() {
		return
	}

	size := heap.list.Size()
	// left_son = index * 2 + 1
	// right_son = left_son + 1
	for si := i<<1 + 1; si < size; si = i<<1 + 1 {
		ri := si + 1

		sv, _ := heap.list.Get(si)
		rv, _ := heap.list.Get(ri)

		lv, li := sv, si

		// check if right is less than left
		if ri < size && heap.cmp(sv, rv) == containers.GreaterThan {
			lv = rv
			li = ri
		}

		// check if iv is less than lv
		iv, _ := heap.list.Get(i)
		if heap.cmp(iv, lv) <= containers.EqualTo {
			break
		}

		lists.Swap[T](heap.list, i, li)
		i = li
	}
}
