package v2

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
	"github.com/jung-kurt/gofpdf"
)

type document struct {
	ctx   context.Context
	_type types.DocumentType
	fpdf  fpdf.Fpdf
	rows  []Row
}

func NewDocument() *document {
	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        "A4",
		FontDirStr:     "",
	})

	fpdf.SetFont("Arial", "B", 16)

	width, height := fpdf.GetPageSize()
	left, top, right, bottom := fpdf.GetMargins()
	fpdf.AddPage()

	return &document{
		fpdf:  fpdf,
		_type: types.Document,
		ctx: context.NewRootContext(width, height, context.Margins{
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
		}),
	}
}

func (d *document) Add(rows ...Row) {
	for _, row := range rows {
		d.rows = append(d.rows, row)
	}
}

func (d *document) Generate(file string) error {
	d.ctx.Print(d._type)

	innerCtx := d.ctx.Copy()
	for _, row := range d.rows {
		row.Render(d.fpdf, innerCtx)
		innerCtx.Coordinate.Y += row.GetHeight()
	}

	return d.fpdf.OutputFileAndClose(file)
}
