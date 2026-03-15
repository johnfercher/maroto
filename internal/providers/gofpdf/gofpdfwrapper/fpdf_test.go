package gofpdfwrapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/phpdave11/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewCustom(t *testing.T) {
	t.Parallel()
	// Act
	sut := gofpdfwrapper.NewCustom(&gofpdf.InitType{})

	// Assert
	assert.Equal(t, "*gofpdf.Fpdf", fmt.Sprintf("%T", sut))
}
