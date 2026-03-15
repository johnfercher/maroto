package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestNewLine(t *testing.T) {
	t.Parallel()
	// Act
	sut := gofpdf.NewLine(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.Line", fmt.Sprintf("%T", sut))
}

func TestLine_Add(t *testing.T) {
	t.Parallel()
	t.Run("when orientation is vertical and color is nil and style is solid, should render vertical line without color/dash calls", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 10, Y: 20, Width: 50, Height: 100}
		prop := &props.Line{
			Orientation:   orientation.Vertical,
			Style:         linestyle.Solid,
			Thickness:     0.5,
			Color:         nil,
			SizePercent:   100,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(5.0, 10.0, 5.0, 10.0)
		fpdf.EXPECT().SetLineWidth(0.5)
		// size = 100 * (100/100) = 100, position = 50 * (50/100) = 25
		// space = (100 - 100) / 2 = 0
		// x = left+cell.X+position = 5+10+25 = 40, y1 = top+cell.Y+space = 10+20+0 = 30, y2 = top+cell.Y+Height-space = 10+20+100-0 = 130
		fpdf.EXPECT().Line(40.0, 30.0, 40.0, 130.0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is vertical and color is set and style is solid, should set and reset draw color", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 200}
		color := &props.Color{Red: 255, Green: 0, Blue: 0}
		prop := &props.Line{
			Orientation:   orientation.Vertical,
			Style:         linestyle.Solid,
			Thickness:     1.0,
			Color:         color,
			SizePercent:   50,
			OffsetPercent: 0,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		fpdf.EXPECT().SetDrawColor(255, 0, 0)
		fpdf.EXPECT().SetLineWidth(1.0)
		// size = 200 * 0.5 = 100, position = 100 * 0 = 0
		// space = (200 - 100) / 2 = 50
		// x = 0+0+0 = 0, y1 = 0+0+50 = 50, y2 = 0+0+200-50 = 150
		fpdf.EXPECT().Line(0.0, 50.0, 0.0, 150.0)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is vertical and style is dashed, should set and reset dash pattern", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 5, Y: 5, Width: 80, Height: 100}
		prop := &props.Line{
			Orientation:   orientation.Vertical,
			Style:         linestyle.Dashed,
			Thickness:     0.2,
			Color:         nil,
			SizePercent:   100,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(10.0, 10.0, 10.0, 10.0)
		fpdf.EXPECT().SetLineWidth(0.2)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		// size = 100*(100/100) = 100, position = 80*(50/100) = 40
		// space = (100-100)/2 = 0
		// x = 10+5+40 = 55, y1 = 10+5+0 = 15, y2 = 10+5+100-0 = 115
		fpdf.EXPECT().Line(55.0, 15.0, 55.0, 115.0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is vertical, color is set and style is dashed, should set color and dash pattern", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 100}
		color := &props.Color{Red: 0, Green: 128, Blue: 255}
		prop := &props.Line{
			Orientation:   orientation.Vertical,
			Style:         linestyle.Dashed,
			Thickness:     0.4,
			Color:         color,
			SizePercent:   50,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		fpdf.EXPECT().SetDrawColor(0, 128, 255)
		fpdf.EXPECT().SetLineWidth(0.4)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		// size = 100*(50/100) = 50, position = 100*(50/100) = 50
		// space = (100-50)/2 = 25
		// x = 0+0+50 = 50, y1 = 0+0+25 = 25, y2 = 0+0+100-25 = 75
		fpdf.EXPECT().Line(50.0, 25.0, 50.0, 75.0)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is horizontal and color is nil and style is solid, should render horizontal line without color/dash calls", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 10, Y: 20, Width: 100, Height: 50}
		prop := &props.Line{
			Orientation:   orientation.Horizontal,
			Style:         linestyle.Solid,
			Thickness:     0.5,
			Color:         nil,
			SizePercent:   100,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(5.0, 10.0, 5.0, 10.0)
		fpdf.EXPECT().SetLineWidth(0.5)
		// size = 100*(100/100) = 100, position = 50*(50/100) = 25
		// space = (100-100)/2 = 0
		// x1 = left+cell.X+space = 5+10+0 = 15, y = top+cell.Y+position = 10+20+25 = 55
		// x2 = left+cell.X+Width-space = 5+10+100-0 = 115
		fpdf.EXPECT().Line(15.0, 55.0, 115.0, 55.0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is horizontal and color is set and style is solid, should set and reset draw color", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 200, Height: 100}
		color := &props.Color{Red: 0, Green: 255, Blue: 0}
		prop := &props.Line{
			Orientation:   orientation.Horizontal,
			Style:         linestyle.Solid,
			Thickness:     1.0,
			Color:         color,
			SizePercent:   50,
			OffsetPercent: 0,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		fpdf.EXPECT().SetDrawColor(0, 255, 0)
		fpdf.EXPECT().SetLineWidth(1.0)
		// size = 200*0.5 = 100, position = 100*0 = 0
		// space = (200-100)/2 = 50
		// x1 = 0+0+50 = 50, y = 0+0+0 = 0, x2 = 0+0+200-50 = 150
		fpdf.EXPECT().Line(50.0, 0.0, 150.0, 0.0)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is horizontal and style is dashed, should set and reset dash pattern", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 5, Y: 5, Width: 100, Height: 80}
		prop := &props.Line{
			Orientation:   orientation.Horizontal,
			Style:         linestyle.Dashed,
			Thickness:     0.2,
			Color:         nil,
			SizePercent:   100,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(10.0, 10.0, 10.0, 10.0)
		fpdf.EXPECT().SetLineWidth(0.2)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		// size = 100*(100/100) = 100, position = 80*(50/100) = 40
		// space = (100-100)/2 = 0
		// x1 = 10+5+0 = 15, y = 10+5+40 = 55, x2 = 10+5+100-0 = 115
		fpdf.EXPECT().Line(15.0, 55.0, 115.0, 55.0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
	t.Run("when orientation is horizontal, color is set and style is dashed, should set color and dash pattern", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 100}
		color := &props.Color{Red: 100, Green: 100, Blue: 100}
		prop := &props.Line{
			Orientation:   orientation.Horizontal,
			Style:         linestyle.Dashed,
			Thickness:     0.3,
			Color:         color,
			SizePercent:   50,
			OffsetPercent: 50,
		}

		fpdf := mocks.NewFpdf(t)
		fpdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		fpdf.EXPECT().SetDrawColor(100, 100, 100)
		fpdf.EXPECT().SetLineWidth(0.3)
		fpdf.EXPECT().SetDashPattern([]float64{1, 1}, 0.0)
		// size = 100*(50/100) = 50, position = 100*(50/100) = 50
		// space = (100-50)/2 = 25
		// x1 = 0+0+25 = 25, y = 0+0+50 = 50, x2 = 0+0+100-25 = 75
		fpdf.EXPECT().Line(25.0, 50.0, 75.0, 50.0)
		fpdf.EXPECT().SetDrawColor(0, 0, 0)
		fpdf.EXPECT().SetLineWidth(linestyle.DefaultLineThickness)
		fpdf.EXPECT().SetDashPattern([]float64{1, 0}, 0.0)

		sut := gofpdf.NewLine(fpdf)

		// Act
		sut.Add(cell, prop)
	})
}
