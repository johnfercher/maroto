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
	text             core.Text
	defaultLineWidth float64
	defaultLineColor props.Color
}

func NewChart(pdf gofpdfwrapper.Fpdf, font core.Font, text core.Text) *chart {
	return &chart{
		pdf:              pdf,
		font:             font,
		text:             text,
		defaultLineWidth: linestyle.DefaultLineThickness,
		defaultLineColor: props.BlackColor,
	}
}

func (s chart) Add(margins *entity.Margins, cell *entity.Cell, width float64, height float64, pps *props.Chart) {
	xPadding := s.getXPadding(pps)
	yPadding := s.getYPadding(pps)

	cellHeight := cell.Height - xPadding
	cellWidth := cell.Width - yPadding

	s.font.SetFont(pps.Scale.Font.Family, pps.Scale.Font.Style, pps.Scale.Font.Size)
	s.pdf.SetDrawColor(s.defaultLineColor.Red, s.defaultLineColor.Green, s.defaultLineColor.Blue)

	stepX, stepY := s.GetSteps(width, height, cellHeight, cellWidth)
	s.horizontalLine(margins, cell, width, stepX, pps)
	s.verticalLine(margins, cell, height, stepY, pps)

	if pps.Title.Text != "" {
		scalePadding := s.font.GetHeight(pps.Scale.Font.Family, pps.Scale.Font.Style, pps.Scale.Font.Size)
		s.text.Add(pps.Title.Text, cell, &props.Text{
			Family: pps.Title.Font.Family,
			Style:  pps.Title.Font.Style,
			Size:   pps.Title.Font.Size,
			Color:  pps.Title.Font.Color,
			Top:    cellHeight + scalePadding,
		})
	}
}

func (s chart) horizontalLine(margins *entity.Margins, cell *entity.Cell, width float64, stepX float64, props *props.Chart) {
	xPadding := s.getXPadding(props)
	yPadding := s.getYPadding(props)

	x := margins.Left + cell.X
	y := cell.Height + margins.Top + cell.Y - xPadding
	widthScale := width * stepX
	s.pdf.Line(x+yPadding, y, x+widthScale+yPadding, y)

	if props.Scale.X == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range props.Scale.X {
		xScaled := label * stepX
		s.pdf.Line(x+xScaled+yPadding, y-1, x+xScaled+yPadding, y+1)

		stringLabel := fmt.Sprintf("%.1f", label)
		stringWidth := s.pdf.GetStringWidth(stringLabel)

		fontFamily, fontType, fontSize := s.font.GetFont()
		stringHeight := s.font.GetHeight(fontFamily, fontType, fontSize)

		s.pdf.Text(x+xScaled-(stringWidth/2.0)+yPadding, y+stringHeight, stringLabel)
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
}

func (s chart) verticalLine(margins *entity.Margins, cell *entity.Cell, height float64, stepY float64, props *props.Chart) {
	//xPadding := s.getXPadding(props)
	//yPadding := s.getYPadding(props)

	x := margins.Left + cell.X
	y := margins.Top + cell.Y
	heightScale := height * stepY
	s.pdf.Line(x, y, x, y+heightScale)

	if props.Scale.Y == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range props.Scale.Y {
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

func (s chart) GetSteps(width, height float64, cellHeight float64, cellWidth float64) (float64, float64) {
	stepX := cellWidth / width
	stepY := cellHeight / height

	return stepX, stepY
}

func (s chart) getXPadding(pps *props.Chart) float64 {
	if pps.Scale.X == nil && pps.Title.Text == "" {
		return 0
	}

	scalePadding := s.font.GetHeight(pps.Scale.Font.Family, pps.Scale.Font.Style, pps.Scale.Font.Size)
	if pps.Title.Text == "" {
		return scalePadding
	}

	titlePadding := s.font.GetHeight(pps.Title.Font.Family, pps.Title.Font.Style, pps.Title.Font.Size)
	return scalePadding + titlePadding
}

func (s chart) getYPadding(pps *props.Chart) float64 {
	if pps.Scale.Y == nil {
		return 0
	}

	return s.font.GetHeight(pps.Scale.Font.Family, pps.Scale.Font.Style, pps.Scale.Font.Size)
}
