package cellwriter_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBorderThicknessStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewBorderThicknessStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderThicknessStyler", fmt.Sprintf("%T", sut))
}
