package lfu

import (
	"github.com/golang-plus/caching"
	"github.com/golang-plus/caching/container/memory"
)

// Container represents a LFU caching container.
type Container struct {
	Capacity int
	heap     *heap
}

func (c *Container) Get(key string) (*caching.Item, error) {
	return c.heap.Get(key), nil
}

func (c *Container) Put(item *caching.Item) error {
	if c.Capacity > 0 && c.heap.Count() == c.Capacity && !c.heap.Contains(item.Key) {
		c.heap.Discard()
	}

	c.heap.Put(item)

	return nil
}

func (c *Container) Remove(key string) error {
	c.heap.Remove(key)

	return nil
}

func (c *Container) Clear() error {
	c.heap.Initialize()

	return nil
}

// NewContainer returns a new in-memory cache container using LFU (least frequently used) arithmetic.
func NewContainer(capacity int) caching.Container {
	return &Container{
		Capacity: capacity,
		heap:     new(heap).Initialize(),
	}
}

// register the container.
func init() {
	memory.LFU.Register(NewContainer)
}
