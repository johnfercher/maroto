package cellwriter

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

type borderThicknessStyler struct {
	StylerTemplate
	defaultBorderSize float64
}

func NewBorderThicknessStyler(fpdf *gofpdf.Fpdf) *borderThicknessStyler {
	return &borderThicknessStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultBorderSize: 0.2,
	}
}

func (f *borderThicknessStyler) Apply(width, height float64, config *config.Config, prop *props.Cell) {
	if prop == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderThickness == 0 {
		f.GoToNext(width, height, config, prop)
		return
	}

	f.fpdf.SetLineWidth(prop.BorderThickness)
	f.GoToNext(width, height, config, prop)
	f.fpdf.SetLineWidth(f.defaultBorderSize)
}
