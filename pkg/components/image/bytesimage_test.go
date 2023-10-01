package image_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFromBytes(t *testing.T) {
	// Act
	sut := image.NewFromBytes(nil, extension.Jpg)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*image.bytesImage", fmt.Sprintf("%T", sut))
}
