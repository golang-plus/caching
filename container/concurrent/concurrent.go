package concurrent

import (
	"sync"

	"github.com/golang-plus/caching"
	"github.com/golang-plus/errors"
)

type Container struct {
	locker sync.RWMutex
	Inner  caching.Container
}

func (c *Container) Clear() error {
	c.locker.Lock()
	defer c.locker.Unlock()

	return c.Inner.Clear()
}

func (c *Container) Remove(key string) error {
	c.locker.Lock()
	defer c.locker.Unlock()

	return c.Inner.Remove(key)
}

func (c *Container) Put(item *caching.Item) error {
	c.locker.Lock()
	defer c.locker.Unlock()

	return c.Inner.Put(item)
}

func (c *Container) Get(key string) (*caching.Item, error) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	return c.Inner.Get(key)
}

// NewContainer returns a new caching container for safe concurrent access.
func NewContainer(inner caching.Container) (caching.Container, error) {
	if inner == nil {
		return nil, errors.New("inner container cannot be nil")
	}

	return &Container{
		Inner: inner,
	}, nil
}

// MustNewContainer is like as NewContainer but panic if inner container is nil.
func MustNewContainer(inner caching.Container) caching.Container {
	c, err := NewContainer(inner)
	if err != nil {
		panic(err)
	}

	return c
}
