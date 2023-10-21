package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
)

func TestNewFillColorStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewFillColorStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.fillColorStyler", fmt.Sprintf("%T", sut))
}
