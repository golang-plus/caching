package arc

import (
	"github.com/golang-plus/caching"
	"github.com/golang-plus/caching/container/memory"
)

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// Container represents a ARC caching container.
type Container struct {
	Capacity int
	p        int // target size of t1
	t1       *list
	t2       *list
	b1       *list
	b2       *list
}

func (c *Container) replace() {
	if c.t1.Count() >= max(1, c.p) { // t1's size exceeds target (t1 is too big)
		c.b1.Put(c.t1.Discard()) // grab from t1 and put to b1
	} else {
		c.b2.Put(c.t2.Discard()) // grab from t2 and put to b2
	}
}

func (c *Container) Get(key string) (*caching.Item, error) {
	if c.t1.Contains(key) { // seen twice recently, put it to t2
		item := c.t1.Remove(key)
		c.t2.Put(item)
		return item, nil
	}

	if c.t2.Contains(key) {
		return c.t2.Get(key).(*caching.Item), nil
	}

	if c.b1.Contains(key) {
		c.p = min(c.Capacity, c.p+max(c.b2.Count()/c.b1.Count(), 1)) // adapt the target size of t1
		c.replace()
		item := c.b1.Remove(key)
		c.t2.Put(item) // seen twice recently, put it to t2
		return item, nil
	}

	if c.b2.Contains(key) {
		c.p = max(0, c.p-max(c.b1.Count()/c.b2.Count(), 1)) // adapt the target size of t1
		c.replace()
		item := c.b2.Remove(key)
		c.t2.Put(item) // seen twice recently, put it to t2
		return item, nil
	}

	return nil, nil
}

func (c *Container) Put(item *caching.Item) error {
	// remove the item anyway
	c.Remove(item.Key)

	if c.t1.Count()+c.b1.Count() == c.Capacity { // b1 + t1 is full
		if c.t1.Count() < c.Capacity { // still room in t1
			c.b1.Discard()
			c.replace()
		} else {
			c.t1.Discard()
		}
	} else { //c.t1.Count()+c.b1.Count() < c.Capacity {
		total := c.t1.Count() + c.t2.Count() + c.b1.Count() + c.b2.Count()
		if total >= c.Capacity { // cache full
			if total == 2*c.Capacity {
				c.b2.Discard()
			}

			c.replace()
		}
	}

	c.t1.Put(item) // seen once recently, put on t1

	return nil
}

func (c *Container) Remove(key string) error {
	c.t1.Remove(key)
	c.t2.Remove(key)
	c.b1.Remove(key)
	c.b2.Remove(key)

	return nil
}

func (c *Container) Clear() error {
	c.p = 0
	c.t1.Init()
	c.t2.Init()
	c.b1.Init()
	c.b2.Init()

	return nil
}

// NewContainer returns a new in-memory cache container using ARC (adaptive/adjustable replacement cache) arithmetic.
func NewContainer(capacity int) caching.Container {
	return &Container{
		Capacity: capacity,
		p:        0,
		t1:       new(list).Initialize(),
		t2:       new(list).Initialize(),
		b1:       new(list).Initialize(),
		b2:       new(list).Initialize(),
	}
}

// register the container.
func init() {
	memory.ARC.Register(NewContainer)
}
