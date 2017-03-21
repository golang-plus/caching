package caching

// Container represents a cache container.
type Container interface {
	// Clear removes all items.
	Clear() error

	// Remove removes the item by given key.
	Remove(key string) error

	// Put inserts/updates the item.
	Put(item *Item) error

	// Get returns the value by given key.
	Get(key string) (*Item, error)
}
