package cache

import (
	"errors"
	"os"
	"sync"
)

type Cache interface {
	GetImage(value string, extension string) (*Image, error)
	LoadImage(value string, extension string) error
	GetCode(code string, codeType string) ([]byte, error)
	SaveCode(code string, codeType string, bytes []byte)
}

type Image struct {
	Bytes     []byte
	Extension string
}

type cache struct {
	images     map[string]*Image
	imageMutex sync.RWMutex
	codes      map[string][]byte
	codeMutex  sync.RWMutex
}

func New() Cache {
	return &cache{
		images:     make(map[string]*Image),
		imageMutex: sync.RWMutex{},
		codes:      make(map[string][]byte),
		codeMutex:  sync.RWMutex{},
	}
}

func (c *cache) LoadImage(file string, extension string) error {
	c.imageMutex.Lock()
	defer c.imageMutex.Unlock()

	imageBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	img := &Image{Bytes: imageBytes, Extension: extension}
	c.images[file+extension] = img

	return nil
}
func (c *cache) GetImage(file string, extension string) (*Image, error) {
	image, ok := c.images[file+extension]
	if ok {
		return image, nil
	}

	return nil, errors.New("image not found")
}

func (c *cache) GetCode(code string, codeType string) ([]byte, error) {
	bytes, ok := c.codes[code+codeType]
	if ok {
		return bytes, nil
	}

	return nil, errors.New("code not found")
}

func (c *cache) SaveCode(code string, codeType string, bytes []byte) {
	c.codeMutex.Lock()
	defer c.codeMutex.Unlock()

	c.codes[code+codeType] = bytes
}
