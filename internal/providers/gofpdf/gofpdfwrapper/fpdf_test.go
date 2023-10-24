package gofpdfwrapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewCustom(t *testing.T) {
	// Act
	sut := gofpdfwrapper.NewCustom(&gofpdf.InitType{})

	// Assert
	assert.NotNil(t, "", fmt.Sprintf("%T", sut))
}
