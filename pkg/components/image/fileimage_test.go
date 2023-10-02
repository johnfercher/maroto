package image_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/stretchr/testify/assert"
)

func TestNewFromFile(t *testing.T) {
	// Act
	sut := image.NewFromFile("path")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*image.fileImage", fmt.Sprintf("%T", sut))
}
