package image_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFromFile(t *testing.T) {
	// Act
	sut := image.NewFromFile("path")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*image.fileImage", fmt.Sprintf("%T", sut))
}
