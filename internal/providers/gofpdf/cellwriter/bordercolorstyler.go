// nolint: dupl
package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type borderColorStyler struct {
	StylerTemplate
	defaultColor *props.Color
}

func NewBorderColorStyler(fpdf gofpdfwrapper.Fpdf) *borderColorStyler {
	return &borderColorStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultColor: &props.BlackColor,
	}
}

func (b *borderColorStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderColor == nil {
		b.GoToNext(width, height, config, prop)
		return
	}

	b.fpdf.SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
	b.GoToNext(width, height, config, prop)
	b.fpdf.SetDrawColor(b.defaultColor.Red, b.defaultColor.Green, b.defaultColor.Blue)
}
