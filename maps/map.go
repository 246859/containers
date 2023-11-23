package maps

import "github.com/246859/containers"

// Map is the base interface of all maps implementations
type Map[K any, V any] interface {
	// Get returns the value corresponding to the given key
	Get(k K) (V, bool)
	// Set sets replace the value corresponding to the given key
	Set(k K, v V)
	// Remove removes the key-value pair from the map
	Remove(k K)
	// Clone returns a clone of the map
	Clone() Map[K, V]
	// Copy copies the src map into the destination map
	Copy(src Map[K, V])
	// Keys returns the all keys of the map
	Keys() []K

	containers.Container[V]
}
