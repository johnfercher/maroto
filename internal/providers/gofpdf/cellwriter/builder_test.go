package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := cellwriter.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.CellWriterBuilder", fmt.Sprintf("%T", sut))
}

func TestCellWriterBuilder_Build(t *testing.T) {
	// Arrange
	sut := cellwriter.NewBuilder()

	// Act
	chain := sut.Build(nil)

	// Assert
	assert.Equal(t, "borderThicknessStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "borderLineStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "borderColorStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "fillColorStyler", chain.GetName())
	chain = chain.GetNext()
	assert.Equal(t, "cellWriter", chain.GetName())
}
