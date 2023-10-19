package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewCellCreator(t *testing.T) {
	// Act
	sut := cellwriter.NewCellCreator(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.cellWriter", fmt.Sprintf("%T", sut))
}
