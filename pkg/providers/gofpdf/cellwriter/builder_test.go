package cellwriter_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf/cellwriter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := cellwriter.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.CellWriterBuilder", fmt.Sprintf("%T", sut))
}
