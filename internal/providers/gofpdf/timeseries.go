package gofpdf

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type timeSeries struct {
	pdf              gofpdfwrapper.Fpdf
	defaultFillColor *props.Color
	defaultDrawColor *props.Color
	defaultLineWidth float64
	chart            core.Chart
	font             core.Font
}

func NewTimeSeries(pdf gofpdfwrapper.Fpdf, chart core.Chart, font core.Font) *timeSeries {
	return &timeSeries{
		pdf:              pdf,
		chart:            chart,
		font:             font,
		defaultFillColor: &props.WhiteColor,
		defaultDrawColor: &props.BlueColor,
		defaultLineWidth: linestyle.DefaultLineThickness,
	}
}

func (s timeSeries) Add(timeSeriesList []entity.TimeSeries, cell *entity.Cell, margins *entity.Margins, prop props.Chart) {
	width, height := s.getSizes(timeSeriesList)
	stepX, stepY := s.getSteps(width, height, cell)

	s.horizontalLine(margins, cell, width, stepX, prop.XLabels)
	s.verticalLine(margins, cell, height, stepY, prop.YLabels)

	s.pdf.SetLineWidth(0.3)

	for _, timeSeries := range timeSeriesList {
		s.pdf.SetDrawColor(timeSeries.Color.Red, timeSeries.Color.Green, timeSeries.Color.Blue)
		for i := 0; i < len(timeSeries.Values)-1; i++ {

			aX := timeSeries.Values[i].X*stepX + margins.Left + cell.X

			aY := timeSeries.Values[i].Y * stepY
			aY = cell.Height + margins.Top + cell.Y - aY

			bX := timeSeries.Values[i+1].X*stepX + margins.Left + cell.X

			bY := timeSeries.Values[i+1].Y * stepY
			bY = cell.Height + margins.Top + cell.Y - bY

			s.pdf.Line(aX, aY, bX, bY)
		}
		s.pdf.SetDrawColor(s.defaultDrawColor.Red, s.defaultDrawColor.Green, s.defaultDrawColor.Blue)
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
}

func (s timeSeries) horizontalLine(margins *entity.Margins, cell *entity.Cell, width float64, stepX float64, labels []float64) {
	x := margins.Left + cell.X
	y := cell.Height + margins.Top + cell.Y
	widthScale := width * stepX
	s.pdf.Line(x, y, x+widthScale, y)

	if labels == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range labels {
		xScaled := label * stepX
		s.pdf.Line(x+xScaled, y-1, x+xScaled, y+1)

		stringLabel := fmt.Sprintf("%.1f", label)
		stringWidth := s.pdf.GetStringWidth(stringLabel)

		s.pdf.Text(x+xScaled-(stringWidth/2.0), y+4, stringLabel)
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
}

func (s timeSeries) verticalLine(margins *entity.Margins, cell *entity.Cell, height float64, stepY float64, labels []float64) {
	x := margins.Left + cell.X
	y := margins.Top + cell.Y
	heightScale := height * stepY
	s.pdf.Line(x, y, x, y+heightScale)

	if labels == nil {
		return
	}

	s.pdf.SetLineWidth(0.3)

	for _, label := range labels {
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

func (s timeSeries) getSizes(timeSeriesList []entity.TimeSeries) (float64, float64) {
	width := 0.0
	height := 0.0

	for _, timeSeries := range timeSeriesList {
		for i := 0; i < len(timeSeries.Values); i++ {
			if timeSeries.Values[i].X > width {
				width = timeSeries.Values[i].X
			}
			if timeSeries.Values[i].Y > height {
				height = timeSeries.Values[i].Y
			}
		}
	}

	return width, height
}

func (s timeSeries) getSteps(width, height float64, cell *entity.Cell) (float64, float64) {
	stepX := cell.Width / width
	stepY := cell.Height / height

	return stepX, stepY
}
