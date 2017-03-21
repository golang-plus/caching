package memory

import (
	"strconv"

	"github.com/golang-plus/caching"
)

// Container represents a memory cache container that implemented in another package.
type Container byte

// Memory Containers.
const (
	// FIFO represents a memory container using replacement algorithm using FIFO (first in first out).
	FIFO Container = iota + 1

	// LFU represents a memory container using replacement algorithm using LFU (least frequently used).
	LFU

	// LRU represents a memory container using replacement algorithm using LRU (least recently used).
	LRU

	// MRU represents a memory container using replacement algorithm using MRU (most recently used).
	MRU

	// ARC represents a memory container using replacement algorithm using ARC (adaptive/adjustable replacement cache).
	ARC

	max
)

var containers = make([]func(int) caching.Container, max)

// Register registers the container container.
// This is intended to be called from the init function in packages that implement container functions.
func (c Container) Register(function func(int) caching.Container) {
	if c <= 0 && c >= max {
		panic("register of unknown memory container function")
	}

	containers[c] = function
}

// Available reports whether the given container is linked into the binary.
func (c Container) Available() bool {
	return c > 0 && c < max && containers[c] != nil
}

// NewContainer returns a new memory container.
func (c Container) NewContainer(capacity int) caching.Container {
	if !c.Available() {
		panic("requested memory container function #" + strconv.Itoa(int(c)) + " is unavailable")
	}

	return containers[c](capacity)
}
