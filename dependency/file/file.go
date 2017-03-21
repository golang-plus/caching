package file

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/golang-plus/caching"
)

// Dependency represents a file dependency.
type Dependency struct {
	Path             string
	LastModifiedTime time.Time
}

// HasChanged reports whether the file has changed.
func (d *Dependency) HasChanged() bool {
	fi, err := os.Stat(d.Path)
	if err != nil {
		return true
	}

	return fi.ModTime().After(d.LastModifiedTime)
}

// NewDependency returns a new file dependency.
func NewDependency(path string) caching.Dependency {
	var lastModifiedTime time.Time
	fi, err := os.Stat(path)
	if err != nil {
		lastModifiedTime = time.Now()
	} else {
		lastModifiedTime = fi.ModTime()
	}

	return &Dependency{
		Path:             path,
		LastModifiedTime: lastModifiedTime,
	}
}

// register for gob.
func init() {
	gob.Register(Dependency{})
}
