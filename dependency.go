package caching

// Dependency represents an external cache dependency.
type Dependency interface {
	// HasChanged reports whehter dependency has changed.
	HasChanged() bool
}
