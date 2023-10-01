package cellwriter

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

type CellWriter interface {
	SetNext(next CellWriter)
	Apply(width, height float64, config *entity.Config, prop *props.Cell)
}

type cellWriter struct {
	StylerTemplate
	defaultColor *props.Color
}

func NewCellCreator(fpdf *gofpdf.Fpdf) *cellWriter {
	return &cellWriter{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultColor: props.NewBlack(),
	}
}

func (c *cellWriter) Apply(width, height float64, config *entity.Config, prop *props.Cell) {
	if prop == nil {
		bd := border.None
		if config.Debug {
			bd = border.Full
		}

		c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", false, 0, "")
		return
	}

	bd := prop.BorderType
	if config.Debug {
		bd = border.Full
	}

	fill := false
	if prop.BackgroundColor != nil {
		c.fpdf.SetFillColor(prop.BackgroundColor.Red, prop.BackgroundColor.Green, prop.BackgroundColor.Blue)
		fill = true
	}

	c.fpdf.CellFormat(width, height, "", string(bd), 0, "C", fill, 0, "")

	if fill {
		white := props.NewWhite()
		c.fpdf.SetFillColor(white.Red, white.Green, white.Blue)
	}
}
