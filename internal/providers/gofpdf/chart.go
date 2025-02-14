package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type chart struct {
	pdf  gofpdfwrapper.Fpdf
	line core.Line
	text core.Text
}

func NewChart(pdf gofpdfwrapper.Fpdf, line core.Line, text core.Text) *chart {
	return &chart{
		pdf:  pdf,
		line: line,
		text: text,
	}
}

func (c *chart) Add(cell *entity.Cell, margins *entity.Margins, prop *props.Chart) {
	// X
	c.line.Add(cell, &props.Line{
		Orientation:   orientation.Horizontal,
		SizePercent:   88,
		OffsetPercent: 94,
		Style:         linestyle.Solid,
		Thickness:     0.5,
	})

	// Y
	c.line.Add(cell, &props.Line{
		Orientation:   orientation.Vertical,
		SizePercent:   88,
		OffsetPercent: 6,
		Style:         linestyle.Solid,
		Thickness:     0.5,
	})
}
