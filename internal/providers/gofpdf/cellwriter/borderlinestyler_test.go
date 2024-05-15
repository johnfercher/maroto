package cellwriter_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"

	"github.com/stretchr/testify/assert"
)

func TestNewBorderLineStyler(t *testing.T) {
	// Act
	sut := cellwriter.NewBorderLineStyler(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cellwriter.borderLineStyler", fmt.Sprintf("%T", sut))
}

func TestBorderLineStyler_Apply(t *testing.T) {
	t.Run("When prop is nil and next is nil, should skip calls", func(t *testing.T) {
		// Arrange
		sut := cellwriter.NewBorderLineStyler(nil)

		// Act
		sut.Apply(100, 100, &entity.Config{}, nil)
	})
	t.Run("When prop is nil and next is filled, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		var nilCellProp *props.Cell

		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, nilCellProp)

		sut := cellwriter.NewBorderLineStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, nilCellProp)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop but line style is solid, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{
			LineStyle: linestyle.Solid,
		}

		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)

		sut := cellwriter.NewBorderLineStyler(nil)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
	})
	t.Run("When has prop but line style is empty, should skip current and call next", func(t *testing.T) {
		// Arrange
		width := 100.0
		height := 100.0
		cfg := &entity.Config{}
		prop := &props.Cell{}

		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)

		sut := cellwriter.NewBorderLineStyler(nil)
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
			LineStyle: linestyle.Dashed,
		}

		inner := mocks.NewCellWriter(t)
		inner.EXPECT().Apply(width, height, cfg, prop)

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)

		sut := cellwriter.NewBorderLineStyler(fpdf)
		sut.SetNext(inner)

		// Act
		sut.Apply(width, height, cfg, prop)

		// Assert
		inner.AssertNumberOfCalls(t, "Apply", 1)
		fpdf.AssertNumberOfCalls(t, "SetDashPattern", 2)
	})
}
