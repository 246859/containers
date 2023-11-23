package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHMap_New(t *testing.T) {
	hmap1 := NewHMap[int, int, string](32, ComparableKeyX[int])
	assert.NotNil(t, hmap1)
}

func TestHMap_Get_Set(t *testing.T) {
	hMap := NewHMap[int, int, string](32, ComparableKeyX[int])
	samples := []struct {
		k int
		v string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}
	for _, sample := range samples {
		hMap.Set(sample.k, sample.v)
	}

	assert.Equal(t, len(samples), hMap.Size())

	for i, sample := range samples {
		v, has := hMap.Get(sample.k)
		assert.True(t, has)
		assert.Equal(t, samples[i].v, v)
	}

	get, b := hMap.Get(100)
	assert.False(t, b)
	assert.Equal(t, "", get)
}

func TestHMap_Clear(t *testing.T) {
	hMap := NewHMap[int, int, string](32, ComparableKeyX[int])
	samples := []struct {
		k int
		v string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}
	for _, sample := range samples {
		hMap.Set(sample.k, sample.v)
	}

	hMap.Clear()
	assert.Equal(t, 0, hMap.Size())
}

func TestHMap_String(t *testing.T) {
	hMap := NewHMap[int, int, string](32, ComparableKeyX[int])
	samples := []struct {
		k int
		v string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}
	for _, sample := range samples {
		hMap.Set(sample.k, sample.v)
	}

	s := hMap.String()
	assert.Greater(t, len(s), 0)
	t.Log(s)
}

func TestHMap_Copy(t *testing.T) {
	hMap := NewHMap[int, int, string](32, ComparableKeyX[int])
	samples := []struct {
		k int
		v string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	}
	for _, sample := range samples {
		hMap.Set(sample.k, sample.v)
	}

	hMap1 := NewHMap[int, int, string](32, ComparableKeyX[int])
	samples1 := []struct {
		k int
		v string
	}{
		{4, "a"},
		{5, "b"},
		{6, "c"},
		{7, "d"},
		{8, "e"},
	}
	for _, sample := range samples1 {
		hMap1.Set(sample.k, sample.v)
	}

	hMap.Copy(hMap1)

	for _, s := range samples1 {
		v, has := hMap.Get(s.k)
		assert.True(t, has)
		assert.Equal(t, s.v, v)
	}
}
