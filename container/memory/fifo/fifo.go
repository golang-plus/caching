package fifo

import (
	"github.com/golang-plus/caching"
	"github.com/golang-plus/caching/container/memory"
)

// Container represents a FIFO caching container.
type Container struct {
	Capacity int
	list     *list
}

func (c *Container) Get(key string) (*caching.Item, error) {
	return c.list.Get(key), nil
}

func (c *Container) Put(item *caching.Item) error {
	if c.Capacity > 0 && c.list.Count() == c.Capacity && !c.list.Contains(item.Key) {
		c.list.Discard()
	}

	c.list.Put(item)

	return nil
}

func (c *Container) Remove(key string) error {
	c.list.Remove(key)

	return nil
}

func (c *Container) Clear() error {
	c.list.Init()

	return nil
}

// NewContainer returns a new in-memory cache container using FIFO (first in first out) arithmetic.
func NewContainer(capacity int) caching.Container {
	return &Container{
		Capacity: capacity,
		list:     new(list).Initialize(),
	}
}

// register the container.
func init() {
	memory.FIFO.Register(NewContainer)
}
