package cellwriter

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

type borderThicknessStyler struct {
	StylerTemplate
	defaultLineThickness float64
}

func NewBorderThicknessStyler(fpdf *gofpdf.Fpdf) *borderThicknessStyler {
	return &borderThicknessStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultLineThickness: linestyle.DefaultLineThickness,
	}
}

func (f *borderThicknessStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
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
	f.fpdf.SetLineWidth(f.defaultLineThickness)
}
