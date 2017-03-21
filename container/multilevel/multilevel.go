package multilevel

import (
	"github.com/golang-plus/caching"
	"github.com/golang-plus/errors"
)

type Container struct {
	List []caching.Container
}

func (c *Container) Clear() error {
	for _, v := range c.List {
		err := v.Clear()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Container) Remove(key string) error {
	for _, v := range c.List {
		err := v.Remove(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Container) Put(item *caching.Item) error {
	for _, v := range c.List {
		err := v.Put(item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Container) Get(key string) (*caching.Item, error) {
	for _, v := range c.List {
		item, err := v.Get(key)
		if err != nil {
			return nil, err
		}

		if item != nil {
			return item, nil
		}
	}

	return nil, nil
}

// NewContainer returns a new multi-level caching container.
func NewContainer(containers ...caching.Container) (caching.Container, error) {
	if len(containers) == 0 {
		return nil, errors.New("containers cannot be empty")
	} else {
		for i, v := range containers {
			if v == nil {
				return nil, errors.Newf("cannot accept nil container [index: %d]", i)
			}
		}
	}

	return &Container{
		List: containers,
	}, nil
}

// MustNewContainer is like as NewContainer but panic if an error happens.
func MustNewContainer(containers ...caching.Container) caching.Container {
	c, err := NewContainer(containers...)
	if err != nil {
		panic(err)
	}

	return c
}
