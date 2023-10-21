package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type borderLineStyler struct {
	StylerTemplate
}

func NewBorderLineStyler(fpdf gofpdfwrapper.Fpdf) *borderLineStyler {
	return &borderLineStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
			name: "borderLineStyler",
		},
	}
}

func (b *borderLineStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.LineStyle == linestyle.Solid || prop.LineStyle == "" {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetDashPattern([]float64{1, 1}, 0)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetDashPattern([]float64{1, 0}, 0)
}
