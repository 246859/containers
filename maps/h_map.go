package maps

import (
	"fmt"
	"maps"
	"strings"
)

// KeyX returns the unique flag of the given key whose type must be comparable
type KeyX[T comparable, K any] func(key K) T

// ComparableKeyX returns it-self
func ComparableKeyX[T comparable](key T) T {
	return key
}

var _ Map[any, any] = (*HMap[uintptr, any, any])(nil)

// NewHMap returns a new hmap whose key can be not-comparable, it according to keyX fn to return the true map key
// if you just want to use a comparable key map, SMap is better.
func NewHMap[T comparable, K any, V any](capacity int, keyFn KeyX[T, K]) *HMap[T, K, V] {
	keys := NewSimpleMap[T, K](capacity)
	values := NewSimpleMap[T, V](capacity)

	return &HMap[T, K, V]{
		keys:   keys,
		values: values,
		keyX:   keyFn,
	}
}

// HMap is implemented by native go map. due to go map key unsupported generic type,
// so we need to use two go map to implement it, keysmap store the key,
// valuesmap store the value, and use a comparable key to related above two maps.
// In this way, the genericization of key is indirectly achieved,
// but it requires more memory to maintain one more go map at runtime.
type HMap[T comparable, K any, V any] struct {
	keys   *SMap[T, K]
	values *SMap[T, V]

	keyX KeyX[T, K]
}

func (hmap *HMap[T, K, V]) Get(k K) (_ V, _ bool) {
	kx := hmap.keyX(k)
	val, has := hmap.values.Get(kx)
	if !has {
		return
	}
	return val, true
}

func (hmap *HMap[T, K, V]) Set(k K, v V) {
	kx := hmap.keyX(k)
	hmap.keys.Set(kx, k)
	hmap.values.Set(kx, v)
}

func (hmap *HMap[T, K, V]) Remove(k K) {
	kx := hmap.keyX(k)
	hmap.keys.Remove(kx)
	hmap.values.Remove(kx)
}

func (hmap *HMap[T, K, V]) Keys() []K {
	return hmap.keys.Values()
}

func (hmap *HMap[T, K, V]) Values() []V {
	return hmap.values.Values()
}

func (hmap *HMap[T, K, V]) Size() int {
	return hmap.keys.Size()
}

func (hmap *HMap[T, K, V]) Clear() {
	hmap.keys.Clear()
	hmap.values.Clear()
}

func (hmap *HMap[T, K, V]) Clone() Map[K, V] {
	ckeys := hmap.keys.Clone().(*SMap[T, K])
	cValues := hmap.values.Clone().(*SMap[T, V])

	return &HMap[T, K, V]{
		keys:   ckeys,
		values: cValues,
		keyX:   hmap.keyX,
	}
}

func (hmap *HMap[T, K, V]) Copy(src Map[K, V]) {
	keys := src.Keys()
	values := src.Values()

	for i, key := range keys {
		kx := hmap.keyX(key)
		hmap.keys.Set(kx, key)
		hmap.values.Set(kx, values[i])
	}
}

func (hmap *HMap[T, K, V]) String() string {
	var strs []string
	for _, kx := range hmap.keys.Keys() {
		k, _ := hmap.keys.Get(kx)
		v, _ := hmap.values.Get(kx)
		strs = append(strs, fmt.Sprintf("%+v:%+v", k, v))
	}
	return fmt.Sprintf("hmap[%s]", strings.Join(strs, ", "))
}

var _ Map[uintptr, any] = (*SMap[uintptr, any])(nil)

func NewSimpleMap[K comparable, V any](capacity int) *SMap[K, V] {
	return &SMap[K, V]{hmap: make(map[K]V, capacity), capacity: capacity}
}

// SMap is simple hmap, implemented by the native go map, just a simple encapsulation
type SMap[K comparable, V any] struct {
	hmap     map[K]V
	capacity int
}

func (hmap *SMap[K, V]) Get(k K) (V, bool) {
	v, has := hmap.hmap[k]
	return v, has
}

func (hmap *SMap[K, V]) Set(k K, v V) {
	hmap.hmap[k] = v
}

func (hmap *SMap[K, V]) Remove(k K) {
	delete(hmap.hmap, k)
}

func (hmap *SMap[K, V]) Clone() Map[K, V] {
	cloneHmap := maps.Clone(hmap.hmap)
	return &SMap[K, V]{
		hmap:     cloneHmap,
		capacity: hmap.capacity,
	}
}

func (hmap *SMap[K, V]) Copy(src Map[K, V]) {
	keys := src.Keys()
	values := src.Values()

	for i, key := range keys {
		if i < len(values) {
			hmap.hmap[key] = values[i]
		}
	}
}

func (hmap *SMap[K, V]) Keys() []K {
	var ks []K
	for k, _ := range hmap.hmap {
		ks = append(ks, k)
	}
	return ks
}

func (hmap *SMap[K, V]) Values() []V {
	var vs []V
	for _, v := range hmap.hmap {
		vs = append(vs, v)
	}
	return vs
}

func (hmap *SMap[K, V]) Size() int {
	return len(hmap.hmap)
}

func (hmap *SMap[K, V]) Clear() {
	clear(hmap.hmap)
}

func (hmap *SMap[K, V]) String() string {
	return fmt.Sprintf("%v", hmap.hmap)
}
