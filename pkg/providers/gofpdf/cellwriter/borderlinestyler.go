package cellwriter

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

type borderLineStyler struct {
	StylerTemplate
}

func NewBorderLineStyler(fpdf *gofpdf.Fpdf) *borderLineStyler {
	return &borderLineStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
	}
}

func (f *borderLineStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	if prop.LineStyle == linestyle.Solid || prop.LineStyle == "" {
		f.GoToNext(width, height, config, prop)
		return
	}

	f.fpdf.SetDashPattern([]float64{1, 1}, 0)
	f.GoToNext(width, height, config, prop)
	f.fpdf.SetDashPattern([]float64{1, 0}, 0)
}
