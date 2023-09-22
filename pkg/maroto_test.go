package pkg_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	// Act
	sut := pkg.NewMaroto()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*pkg.maroto", fmt.Sprintf("%T", sut))
}
