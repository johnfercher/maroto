package cache_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMutexDecorator(t *testing.T) {
	// Act
	sut := cache.NewMutexDecorator(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.mutexCache", fmt.Sprintf("%T", sut))
}