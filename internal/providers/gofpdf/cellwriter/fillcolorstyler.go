// nolint: dupl
package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type fillColorStyler struct {
	StylerTemplate
	defaultFillColor *props.Color
}

func NewFillColorStyler(fpdf gofpdfwrapper.Fpdf) *fillColorStyler {
	return &fillColorStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultFillColor: &props.WhiteColor,
	}
}

func (f *fillColorStyler) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	if prop.BackgroundColor == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	f.fpdf.SetFillColor(prop.BackgroundColor.Red, prop.BackgroundColor.Green, prop.BackgroundColor.Blue)
	f.GoToNext(width, height, config, prop)
	f.fpdf.SetFillColor(f.defaultFillColor.Red, f.defaultFillColor.Green, f.defaultFillColor.Blue)
}
