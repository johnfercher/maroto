package chart

import (
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
	chart            core.Chart
	font             core.Font
	defaultLineWidth float64
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

func (s timeSeries) Add(timeSeriesList []entity.TimeSeries, cell *entity.Cell, margins *entity.Margins, prop *props.Chart) {
	width, height := s.getSizes(timeSeriesList)

	stepX, stepY := s.chart.GetSteps(width, height, cell.Height, cell.Width)

	s.pdf.SetLineWidth(0.3)

	for _, timeSeries := range timeSeriesList {
		s.pdf.SetDrawColor(timeSeries.Color.Red, timeSeries.Color.Green, timeSeries.Color.Blue)

		for i := 0; i < len(timeSeries.Points)-1; i++ {
			aX := timeSeries.Points[i].X*stepX + margins.Left + cell.X

			aY := timeSeries.Points[i].Y * stepY
			aY = cell.Height + margins.Top + cell.Y - aY

			bX := timeSeries.Points[i+1].X*stepX + margins.Left + cell.X

			bY := timeSeries.Points[i+1].Y * stepY
			bY = cell.Height + margins.Top + cell.Y - bY

			s.pdf.Line(aX, aY, bX, bY)
		}

		for i := 0; i < len(timeSeries.Labels); i++ {
			aX := timeSeries.Labels[i].Point.X*stepX + margins.Left + cell.X

			aY := timeSeries.Labels[i].Point.Y * stepY
			aY = cell.Height + margins.Top + cell.Y - aY

			stringLabel := timeSeries.Labels[i].Value
			stringWidth := s.pdf.GetStringWidth(stringLabel)

			fontFamily, fontType, fontSize := s.font.GetFont()
			stringHeight := s.font.GetHeight(fontFamily, fontType, fontSize)

			s.pdf.Circle(aX, aY, 0.6, "D")
			s.pdf.Text(aX-(stringWidth/2), aY-(stringHeight/1.5), stringLabel)
		}

		s.pdf.SetDrawColor(s.defaultDrawColor.Red, s.defaultDrawColor.Green, s.defaultDrawColor.Blue)
	}

	s.pdf.SetLineWidth(s.defaultLineWidth)
	s.chart.Add(margins, cell, width, height, prop)
}

func (s timeSeries) getSizes(timeSeriesList []entity.TimeSeries) (float64, float64) {
	width := 0.0
	height := 0.0

	for _, timeSeries := range timeSeriesList {
		for i := 0; i < len(timeSeries.Points); i++ {
			if timeSeries.Points[i].X > width {
				width = timeSeries.Points[i].X
			}
			if timeSeries.Points[i].Y > height {
				height = timeSeries.Points[i].Y
			}
		}
	}

	return width, height
}
