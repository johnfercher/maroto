package cache

import (
	"sync"
)

type mutexCache struct {
	inner      Cache
	imageMutex sync.RWMutex
	codeMutex  sync.RWMutex
}

func NewMutexDecorator(cache Cache) Cache {
	return &mutexCache{
		inner:      cache,
		imageMutex: sync.RWMutex{},
		codeMutex:  sync.RWMutex{},
	}
}

func (c *mutexCache) LoadImage(file string, extension string) error {
	c.imageMutex.Lock()
	defer c.imageMutex.Unlock()

	return c.inner.LoadImage(file, extension)
}

func (c *mutexCache) GetImage(file string, extension string) (*Image, error) {
	return c.inner.GetImage(file, extension)
}

func (c *mutexCache) GetCode(code string, codeType string) ([]byte, error) {
	return c.inner.GetCode(code, codeType)
}

func (c *mutexCache) SaveCode(code string, codeType string, bytes []byte) {
	c.codeMutex.Lock()
	defer c.codeMutex.Unlock()

	c.inner.SaveCode(code, codeType, bytes)
}
