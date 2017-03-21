package caching

import (
	"github.com/golang-plus/errors"
)

// Cache represents a cache manager.
type Cache struct {
	container Container
}

// SetContainer sets the container for cache.
func (c *Cache) SetContainer(v Container) {
	c.container = v
}

// Clear removes all items from cache.
func (c *Cache) Clear() error {
	err := c.container.Clear()
	if err != nil {
		return errors.Wrap(err, "could not clear cache items")
	}

	return nil
}

// Remove removes the cache item by given key.
func (c *Cache) Remove(key string) error {
	if len(key) == 0 {
		return errors.New("key of item cannot be empty")
	}

	err := c.container.Remove(key)
	if err != nil {
		return errors.Wrapf(err, "could not remove item with key %q", key)
	}

	return nil
}

// Put saves the cache item.
func (c *Cache) Put(item *Item) error {
	if item == nil {
		return errors.New("item cannot be nil")
	}

	err := c.container.Put(item)
	if err != nil {
		return errors.Wrapf(err, "could not put cache item with key %q to container", item.Key)
	}

	return nil
}

// Get returns the cache item by given key.
// It returns nil if cache item has expired or not found.
func (c *Cache) Get(key string) (*Item, error) {
	if len(key) == 0 {
		return nil, errors.New("key of item cannot be empty")
	}

	item, err := c.container.Get(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get item with key %q", key)
	}
	if item == nil {
		return nil, nil
	}
	if item.HasExpired() {
		err := c.container.Remove(key)
		if err != nil {
			return nil, errors.Wrapf(err, "could not remove expired item with key %q", key)
		}

		return nil, nil
	}
	item.touch()                // update last accessed time
	err = c.container.Put(item) // save the item to container
	if err != nil {
		return nil, errors.Wrapf(err, "could not update item with key %q", item.Key)
	}

	return item, nil
}

// NewCache returns a new cache.
func NewCache(container Container) (*Cache, error) {
	if container == nil {
		return nil, errors.New("container of cache cannot be nil")
	}

	return &Cache{
		container: container,
	}, nil
}

// MustNewCache is like as NewCache but panic if container is nil.
func MustNewCache(container Container) *Cache {
	c, err := NewCache(container)
	if err != nil {
		panic(err)
	}

	return c
}
