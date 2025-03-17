package gofpdf

import (
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type timeSeries struct {
	pdf              gofpdfwrapper.Fpdf
	defaultFillColor *props.Color
	chart            core.Chart
}

func NewTimeSeries(pdf gofpdfwrapper.Fpdf, chart core.Chart) *timeSeries {
	return &timeSeries{
		pdf:              pdf,
		chart:            chart,
		defaultFillColor: &props.WhiteColor,
	}
}

func (s timeSeries) Add(timeSeriesList []entity.TimeSeries, cell *entity.Cell, margins *entity.Margins) {
	stepX, stepY := s.getSteps(timeSeriesList, cell)

	for _, timeSeries := range timeSeriesList {
		for i := 0; i < len(timeSeries.Values)-1; i++ {
			s.pdf.SetDrawColor(timeSeries.Color.Red, timeSeries.Color.Green, timeSeries.Color.Blue)

			aX := timeSeries.Values[i].X*stepX + margins.Left + cell.X
			aY := timeSeries.Values[i].Y*stepY + margins.Top + cell.Y
			bX := timeSeries.Values[i+1].X*stepX + margins.Left + cell.X
			bY := timeSeries.Values[i+1].Y*stepY + margins.Top + cell.Y

			s.pdf.Line(aX, aY, bX, bY)

			s.pdf.SetDrawColor(0, 0, 0)
		}
	}
}

func (s timeSeries) getSteps(timeSeriesList []entity.TimeSeries, cell *entity.Cell) (float64, float64) {
	xSize := 0.0
	ySize := 0.0

	for _, timeSeries := range timeSeriesList {
		for i := 0; i < len(timeSeries.Values); i++ {
			if timeSeries.Values[i].X > xSize {
				xSize = timeSeries.Values[i].X
			}
			if timeSeries.Values[i].Y > ySize {
				ySize = timeSeries.Values[i].Y
			}
		}
	}

	stepX := cell.Width / xSize
	stepY := cell.Height / ySize

	return stepX, stepY
}
