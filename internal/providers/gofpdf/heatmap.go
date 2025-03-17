package gofpdf

import (
	"errors"
	"math"

	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

var ErrOutOfRange = errors.New("out of range")

type heatMap struct {
	pdf              gofpdfwrapper.Fpdf
	defaultFillColor *props.Color
	chart            core.Chart
	padding          float64
}

func NewHeatMap(pdf gofpdfwrapper.Fpdf, chart core.Chart) *heatMap {
	return &heatMap{
		pdf:              pdf,
		chart:            chart,
		defaultFillColor: &props.WhiteColor,
		padding:          0,
	}
}

func (s heatMap) Add(heatMap [][]int, cell *entity.Cell, margins *entity.Margins, prop *props.HeatMap) {
	if heatMap == nil || len(heatMap) == 0 || len(heatMap[0]) == 0 {
		return
	}

	max := s.getMax(heatMap)
	transparent := s.getTransparent(prop)
	stepX, stepY := s.getSteps(heatMap, cell)

	for i := 0; i < len(heatMap)-1; i++ {
		for j := 0; j < len(heatMap[i])-1; j++ {
			if !transparent[heatMap[i][j]] {
				r, g, b := GetHeatColor(heatMap[i][j], max, prop.InvertScale, prop.HalfColor)

				x := float64(i)*stepX + cell.X + margins.Left
				y := float64(j) * stepY

				// Invert to draw from bottom to up
				y = cell.Height + margins.Top + cell.Y - y - stepY

				s.pdf.SetFillColor(r, g, b)
				s.pdf.Rect(x, y, stepX, stepY, "F")
				s.pdf.SetFillColor(s.defaultFillColor.Red, s.defaultFillColor.Green, s.defaultFillColor.Blue)
			}
		}
	}

	if prop.Chart != nil {
		s.chart.Add(cell, margins, prop.Chart)
	}
}

func (s heatMap) getSteps(heatMap [][]int, cell *entity.Cell) (float64, float64) {
	xSize := len(heatMap)
	stepX := (cell.Width) / float64(xSize-1)

	ySize := len(heatMap[0])
	stepY := (cell.Height) / float64(ySize-1)

	return stepX, stepY
}

func GetHeatColor(i int, total int, invertScale bool, halfColor bool) (int, int, int) {
	offset := 360.0 / 7.0 / 2.0
	iStep := GetStepWithOffset(360, float64(total), float64(i), -offset)

	if iStep < 0 {
		iStep = 360 + iStep
	}

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

func GetStepWithOffset(scaleMax float64, valueMax float64, i float64, offset float64) float64 {
	scaleStep := GetStep(scaleMax, valueMax)
	iStep := i * scaleStep
	return iStep + offset
}
