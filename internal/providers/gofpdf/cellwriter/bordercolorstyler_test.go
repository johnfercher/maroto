package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewBorderColorStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewBorderColorStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderColorStyler", fmt.Sprintf("%T", sut))
}
