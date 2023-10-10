package internal

import (
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type line struct {
	pdf              fpdf.Fpdf
	defaultColor     *props.Color
	defaultThickness float64
}

func NewLine(pdf fpdf.Fpdf) *line {
	return &line{
		pdf:              pdf,
		defaultColor:     &props.BlackColor,
		defaultThickness: linestyle.DefaultLineThickness,
	}
}

func (l *line) Add(cell *entity.Cell, prop *props.Line) {
	if prop.Orientation == orientation.Vertical {
		l.renderVertical(cell, prop)
	} else {
		l.renderHorizontal(cell, prop)
	}
}

func (l *line) renderVertical(cell *entity.Cell, prop *props.Line) {
	size := cell.Height * (prop.SizePercent / 100.0)
	position := cell.Width * (prop.OffsetPercent / 100.0)

	space := (cell.Height - size) / 2.0

	left, top, _, _ := l.pdf.GetMargins()

	if prop.Color != nil {
		l.pdf.SetDrawColor(prop.Color.Red, prop.Color.Green, prop.Color.Blue)
	}
	l.pdf.SetLineWidth(prop.Thickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 1}, 0)
	}

	l.pdf.Line(left+cell.X+position, top+cell.Y+space, left+cell.X+position, top+cell.Y+cell.Height-space)

	if prop.Color != nil {
		l.pdf.SetDrawColor(l.defaultColor.Red, l.defaultColor.Green, l.defaultColor.Blue)
	}
	l.pdf.SetLineWidth(l.defaultThickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 0}, 0)
	}
}

func (l *line) renderHorizontal(cell *entity.Cell, prop *props.Line) {
	size := cell.Width * (prop.SizePercent / 100.0)
	position := cell.Height * (prop.OffsetPercent / 100.0)

	space := (cell.Width - size) / 2.0

	left, top, _, _ := l.pdf.GetMargins()

	if prop.Color != nil {
		l.pdf.SetDrawColor(prop.Color.Red, prop.Color.Green, prop.Color.Blue)
	}
	l.pdf.SetLineWidth(prop.Thickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 1}, 0)
	}

	l.pdf.Line(left+cell.X+space, top+cell.Y+position, left+cell.X+cell.Width-space, top+cell.Y+position)

	if prop.Color != nil {
		l.pdf.SetDrawColor(l.defaultColor.Red, l.defaultColor.Green, l.defaultColor.Blue)
	}
	l.pdf.SetLineWidth(l.defaultThickness)

	if prop.Style != linestyle.Solid {
		l.pdf.SetDashPattern([]float64{1, 0}, 0)
	}
}
