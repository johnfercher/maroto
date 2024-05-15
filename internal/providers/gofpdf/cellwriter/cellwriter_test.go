package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewCellCreator(t *testing.T) {
	// Act
	sut := cellwriter.NewCellWriter(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.cellWriter", fmt.Sprintf("%T", sut))
}

func TestCellWriter_Apply(t *testing.T) {
	t.Run("when prop is nil without debug, should call cellformat correctly", func(t *testing.T) {
		// Arrange
		config := &entity.Config{}
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "", 0, "C", false, 0, "")

		sut := cellwriter.NewCellWriter(fpdf)

		// Act
		sut.Apply(width, height, config, nil)

		// Assert
		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})
	t.Run("when prop is nil with debug, should call cellformat correctly", func(t *testing.T) {
		// Arrange
		config := &entity.Config{
			Debug: true,
		}
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "1", 0, "C", false, 0, "")

		sut := cellwriter.NewCellWriter(fpdf)

		// Act
		sut.Apply(width, height, config, nil)

		// Assert
		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})
	t.Run("when has prop without debug, should call cellformat correctly", func(t *testing.T) {
		// Arrange
		config := &entity.Config{}
		prop := fixture.CellProp()
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "L", 0, "C", true, 0, "")

		sut := cellwriter.NewCellWriter(fpdf)

		// Act
		sut.Apply(width, height, config, &prop)

		// Assert
		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})
	t.Run("when has prop with debug, should call cellformat correctly", func(t *testing.T) {
		// Arrange
		config := &entity.Config{
			Debug: true,
		}
		prop := fixture.CellProp()
		width := 100.0
		height := 200.0
		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().CellFormat(width, height, "", "1", 0, "C", true, 0, "")

		sut := cellwriter.NewCellWriter(fpdf)

		// Act
		sut.Apply(width, height, config, &prop)

		// Assert
		fpdf.AssertNumberOfCalls(t, "CellFormat", 1)
	})
}
