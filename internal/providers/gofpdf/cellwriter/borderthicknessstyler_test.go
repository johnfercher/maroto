package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewBorderThicknessStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewBorderThicknessStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderThicknessStyler", fmt.Sprintf("%T", sut))
}
