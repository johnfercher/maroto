package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"

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

func TestBorderColorStyler_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		// Arrange
		sut := cellwriter.NewBorderColorStyler(nil)

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

		sut := cellwriter.NewBorderColorStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, nilCellProp)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop but border color is nil, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{}

		inner := &mocks.CellWriter{}
		inner.EXPECT().Apply(width, height, cfg, prop)

		sut := cellwriter.NewBorderColorStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop and border color is defined, should apply current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{
			BorderColor: &props.Color{Red: 140, Green: 100, Blue: 80},
		}

		inner := &mocks.CellWriter{}
		inner.EXPECT().Apply(width, height, cfg, prop)

		fpdf := &mocks.Fpdf{}
		fpdf.EXPECT().SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)

		sut := cellwriter.NewBorderColorStyler(fpdf)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
	})
}
