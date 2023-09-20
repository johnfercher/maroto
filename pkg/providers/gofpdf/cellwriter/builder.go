package cellwriter

import "github.com/jung-kurt/gofpdf"

type CellWriterBuilder struct{}

func NewBuilder() *CellWriterBuilder {
	return &CellWriterBuilder{}
}

func (c *CellWriterBuilder) Build(fpdf *gofpdf.Fpdf) CellWriter {
	cellCreator := NewCellCreator(fpdf)
	borderColorStyle := NewBorderColorStyler(fpdf)
	borderLineStyler := NewBorderLineStyler(fpdf)
	borderThicknessStyler := NewBorderThicknessStyler(fpdf)

	borderThicknessStyler.SetNext(borderLineStyler)
	borderLineStyler.SetNext(borderColorStyle)
	borderColorStyle.SetNext(cellCreator)

	return borderThicknessStyler
}
