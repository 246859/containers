package lists

import (
	"github.com/246859/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	var list List[int]
	list = NewArrayList[int](32, containers.OrderedEqual[int])

	list.Add(1, 2, 3, 4, 5)

	Swap(list, 1, 2)

	get1, b1 := list.Get(1)
	assert.True(t, b1)
	assert.Equal(t, 3, get1)
	get2, b2 := list.Get(2)
	assert.True(t, b2)
	assert.Equal(t, 2, get2)
}
