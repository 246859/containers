package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	it := NewSliceIterator[int](data)
	for it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}
	// reverse
	it.Reverse()
	it.Reverse()
	for it.Next() {
		assert.Equal(t, data[it.Index()], it.Value())
	}
}
