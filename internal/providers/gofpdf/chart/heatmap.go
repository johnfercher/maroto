package chart

import (
	"errors"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"math"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var ErrOutOfRange = errors.New("out of range")

type heatMap struct {
	pdf              gofpdfwrapper.Fpdf
	defaultFillColor *props.Color
	chart            core.Chart
	font             core.Font
	padding          float64
}

func NewHeatMap(pdf gofpdfwrapper.Fpdf, chart core.Chart, font core.Font) *heatMap {
	return &heatMap{
		pdf:              pdf,
		chart:            chart,
		font:             font,
		defaultFillColor: &props.WhiteColor,
		padding:          0,
	}
}

func (s heatMap) Add(heatMap [][]int, cell *entity.Cell, margins *entity.Margins, prop *props.HeatMap) {
	xPadding := s.getXPadding(prop)
	yPadding := s.getYPadding(prop)

	cellHeight := cell.Height - xPadding
	cellWidth := cell.Width - yPadding

	if heatMap == nil || len(heatMap) == 0 || len(heatMap[0]) == 0 {
		return
	}

	max := s.getMax(heatMap)
	width := float64(len(heatMap))
	height := float64(len(heatMap[0]))
	transparent := s.getTransparent(prop)
	stepX, stepY := s.chart.GetSteps(width, height, cellHeight, cellWidth)

	for i := 0; i < len(heatMap); i++ {
		for j := 0; j < len(heatMap[i]); j++ {
			currentValue := heatMap[i][j]
			_, isTransparent := transparent[currentValue]
			if isTransparent {
				continue
			}

			r, g, b := GetHeatColor(currentValue, max, prop.HSVScale)

			x := float64(i)*stepX + cell.X + margins.Left + yPadding
			y := float64(j)*stepY + xPadding

			// Invert to draw from bottom to up
			y = cell.Height + margins.Top + cell.Y - y - stepY

			s.pdf.SetFillColor(r, g, b)
			s.pdf.Rect(x, y, stepX, stepY, "F")
			s.pdf.SetFillColor(s.defaultFillColor.Red, s.defaultFillColor.Green, s.defaultFillColor.Blue)
		}
	}

	s.chart.Add(margins, cell, width, height, &prop.Chart)
}

func (s heatMap) getXPadding(pps *props.HeatMap) float64 {
	if pps.Chart.Scale.X == nil && pps.Chart.Title.Text == "" {
		return 0
	}

	scalePadding := s.font.GetHeight(pps.Chart.Scale.Font.Family, pps.Chart.Scale.Font.Style, pps.Chart.Scale.Font.Size)
	if pps.Chart.Title.Text == "" {
		return scalePadding
	}

	titlePadding := s.font.GetHeight(pps.Chart.Title.Font.Family, pps.Chart.Title.Font.Style, pps.Chart.Title.Font.Size)
	return scalePadding + titlePadding
}

func (s heatMap) getYPadding(pps *props.HeatMap) float64 {
	if pps.Chart.Scale.Y == nil {
		return 0
	}

	return s.font.GetHeight(pps.Chart.Scale.Font.Family, pps.Chart.Scale.Font.Style, pps.Chart.Scale.Font.Size)
}

func GetHeatColor(i int, total int, scale props.HSVScale) (int, int, int) {
	iStep := GetStepWithOffset(float64(scale.End)-float64(scale.Begin), float64(total), float64(i))

	iStep += float64(scale.Begin)

	r, g, b, _ := HSVToRGB(iStep, 1.0, 1.0)
	return int(r), int(g), int(b)
}

func (s heatMap) getMax(matrix [][]int) int {
	var max = 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell > max {
				max = cell
			}
		}
	}

	return max
}

func (s heatMap) getTransparent(p *props.HeatMap) map[int]bool {
	m := make(map[int]bool)
	for _, t := range p.TransparentValues {
		m[t] = true
	}
	return m
}

// HSVToRGB converts an HSV triple to an RGB triple.
// Source: https://github.com/Crazy3lf/colorconv/blob/master/colorconv.go
func HSVToRGB(h, s, v float64) (r, g, b uint8, err error) {
	if h < 0 || h >= 360 ||
		s < 0 || s > 1 ||
		v < 0 || v > 1 {
		return 0, 0, 0, ErrOutOfRange
	}
	// When 0 ≤ h < 360, 0 ≤ s ≤ 1 and 0 ≤ v ≤ 1:
	C := v * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := v - C
	var Rnot, Gnot, Bnot float64
	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}
	r = uint8(math.Round((Rnot + m) * 255))
	g = uint8(math.Round((Gnot + m) * 255))
	b = uint8(math.Round((Bnot + m) * 255))
	return r, g, b, nil
}

func GetStep(scaleMax float64, valueMax float64) float64 {
	return scaleMax / valueMax
}

func GetStepWithOffset(scaleMax float64, valueMax float64, i float64) float64 {
	scaleStep := GetStep(scaleMax, valueMax)
	iStep := i * scaleStep
	return iStep
}
