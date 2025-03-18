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
	defaultLineWidth float64
}

func NewTimeSeries(pdf gofpdfwrapper.Fpdf, chart core.Chart) *timeSeries {
	return &timeSeries{
		pdf:              pdf,
		chart:            chart,
		defaultFillColor: &props.WhiteColor,
		defaultDrawColor: &props.BlueColor,
		defaultLineWidth: linestyle.DefaultLineThickness,
	}
}

func (s timeSeries) Add(timeSeriesList []entity.TimeSeries, cell *entity.Cell, margins *entity.Margins, prop *props.Chart) {
	width, height := s.getSizes(timeSeriesList)

	stepX, stepY := s.chart.GetSteps(width, height, cell)

	s.chart.Add(margins, cell, width, height, prop)

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
