# caching

Package **caching** providers a scalable cache component.

## Container

The container represents an adapter for cache manager/service.

### Builtin Containers

* Concurrent Container: wrapping a container for concurrent access.
* Multi-Level Container: wrapping the containers into a single container.
* Memory Containers: local memory containers (not safe for concurrent access).
    * FIFO Container: replacement algorithm using FIFO (first in first out).
    * LFU Container: replacement algorithm using LFU (least frequently used).
    * LRU Container: replacement algorithm using LRU (least recently used).
    * MRU Container: replacement algorithm using MRU (most recently used).
    * ARC Container: replacement algorithm using ARC (adaptive/adjustable replacement cache).

## Dependency

The dependency represents an external expiration policy.

### Builtin Dependencies

* File Dependency: it's useful for caching configuration file.
