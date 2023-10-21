package cellwriter_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
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

func TestFillColorStyle_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		// Arrange
		sut := cellwriter.NewFillColorStyler(nil)

		// Act
		sut.Apply(100, 100, &entity.Config{}, nil)
	})
	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		var nilCellProp *props.Cell

		inner := &mocks.CellWriter{}
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)

		sut := cellwriter.NewFillColorStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, nilCellProp)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop but background color is nil, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{}

		inner := &mocks.CellWriter{}
		inner.EXPECT().Apply(width, height, cfg, prop)

		sut := cellwriter.NewFillColorStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop and line style is dashed, should apply current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{
			BorderThickness: 1.0,
		}

		inner := &mocks.CellWriter{}
		inner.EXPECT().Apply(width, height, cfg, prop)

		fpdf := &mocks.Fpdf{}
		fpdf.EXPECT().SetFillColor(prop.BorderThickness)
		fpdf.EXPECT().SetFillColor(linestyle.DefaultLineThickness)

		sut := cellwriter.NewBorderThicknessStyler(fpdf)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)
	})
}
