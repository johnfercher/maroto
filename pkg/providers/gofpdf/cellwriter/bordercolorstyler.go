package cellwriter

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

type borderColorStyler struct {
	StylerTemplate
	defaultColor *props.Color
}

func NewBorderColorStyler(fpdf *gofpdf.Fpdf) *borderColorStyler {
	return &borderColorStyler{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultColor: props.NewBlack(),
	}
}

func (f *borderColorStyler) Apply(width, height float64, config *config.Config, prop *props.Cell) {
	if prop == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	if prop.BorderColor == nil {
		f.GoToNext(width, height, config, prop)
		return
	}

	f.fpdf.SetDrawColor(prop.BorderColor.Red, prop.BorderColor.Green, prop.BorderColor.Blue)
	f.GoToNext(width, height, config, prop)
	f.fpdf.SetDrawColor(f.defaultColor.Red, f.defaultColor.Green, f.defaultColor.Blue)
}
