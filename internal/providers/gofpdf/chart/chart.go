package chart

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type chart struct {
	pdf              gofpdfwrapper.Fpdf
	font             core.Font
	defaultLineWidth float64
}

func NewChart(pdf gofpdfwrapper.Fpdf, font core.Font) *chart {
	return &chart{
		pdf:              pdf,
		font:             font,
		defaultLineWidth: linestyle.DefaultLineThickness,
	}
}

func (s chart) Add(margins *entity.Margins, cell *entity.Cell, width float64, height float64, props *props.Chart) {
	if props.Font != nil {
		s.font.SetFont(props.Font.Family, props.Font.Style, props.Font.Size)
	}

	stepX, stepY := s.GetSteps(width, height, cell)
	s.horizontalLine(margins, cell, width, stepX, props)
	s.verticalLine(margins, cell, height, stepY, props)
}

func (s chart) horizontalLine(margins *entity.Margins, cell *entity.Cell, width float64, stepX float64, props *props.Chart) {
	x := margins.Left + cell.X
	y := cell.Height + margins.Top + cell.Y
	widthScale := width * stepX
	s.pdf.Line(x, y, x+widthScale, y)

	if props.XLabels == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range props.XLabels {
		xScaled := label * stepX
		s.pdf.Line(x+xScaled, y-1, x+xScaled, y+1)

		stringLabel := fmt.Sprintf("%.1f", label)
		stringWidth := s.pdf.GetStringWidth(stringLabel)

		fontFamily, fontType, fontSize := s.font.GetFont()
		stringHeight := s.font.GetHeight(fontFamily, fontType, fontSize)

		s.pdf.Text(x+xScaled-(stringWidth/2.0), y+stringHeight, stringLabel)
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
}

func (s chart) verticalLine(margins *entity.Margins, cell *entity.Cell, height float64, stepY float64, props *props.Chart) {
	x := margins.Left + cell.X
	y := margins.Top + cell.Y
	heightScale := height * stepY
	s.pdf.Line(x, y, x, y+heightScale)

	if props.YLabels == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range props.YLabels {
		yScaled := label * stepY
		s.pdf.Line(x-1, heightScale-yScaled+y, x+1, heightScale-yScaled+y)

		stringLabel := fmt.Sprintf("%.2f", label)
		stringWidth := s.pdf.GetStringWidth(stringLabel)

		fontFamily, fontType, fontSize := s.font.GetFont()
		stringHeight := s.font.GetHeight(fontFamily, fontType, fontSize)

		s.pdf.Text(x-stringWidth, heightScale-yScaled+y+(stringHeight/2.0)-0.5, fmt.Sprintf("%.1f", label))
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
}

func (s chart) GetSteps(width, height float64, cell *entity.Cell) (float64, float64) {
	stepX := cell.Width / width
	stepY := cell.Height / height

	return stepX, stepY
}
