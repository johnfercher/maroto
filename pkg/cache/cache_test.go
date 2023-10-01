package cache_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := cache.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.cache", fmt.Sprintf("%T", sut))
}
