package cache

import (
	"sync"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type mutexCache struct {
	inner      Cache
	imageMutex sync.RWMutex
}

// NewMutexDecorator is responsible to create a new Cache.
func NewMutexDecorator(cache Cache) Cache {
	return &mutexCache{
		inner:      cache,
		imageMutex: sync.RWMutex{},
	}
}

// LoadImage loads an image from a file.
func (c *mutexCache) LoadImage(file string, extension extension.Type) error {
	c.imageMutex.Lock()
	defer c.imageMutex.Unlock()

	return c.inner.LoadImage(file, extension)
}

// AddImage adds an image to the cache.
func (c *mutexCache) AddImage(value string, image *entity.Image) {
	c.imageMutex.Lock()
	defer c.imageMutex.Unlock()

	c.inner.AddImage(value, image)
}

// GetImage returns an image from the cache.
func (c *mutexCache) GetImage(file string, extension extension.Type) (*entity.Image, error) {
	return c.inner.GetImage(file, extension)
}
