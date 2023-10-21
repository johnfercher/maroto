package linewriter

/*type LineWriter interface {
	SetNext(next LineWriter)
	Apply(cell *entity.Cell, prop *props.Line)
}

type lineWriter struct {
	StylerTemplate
	defaultColor *props.Color
}

func NewCellCreator(fpdf gofpdf.Fpdf) *lineWriter {
	return &lineWriter{
		StylerTemplate: StylerTemplate{
			fpdf: fpdf,
		},
		defaultColor: &props.BlackColor,
	}
}

func (c *lineWriter) Apply(cell *entity.Cell, prop *props.Line) {
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
		white := &props.WhiteColor
		c.fpdf.SetFillColor(white.Red, white.Green, white.Blue)
	}
}
*/
