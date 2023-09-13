package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
	"github.com/jung-kurt/gofpdf"
)

type document struct {
	file  string
	ctx   context.Context
	_type types.DocumentType
	fpdf  fpdf.Fpdf
	pages []Page
	rows  []Row
}

func NewDocument(file string) *document {
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
		file:  file,
		fpdf:  fpdf,
		_type: types.Document,
		ctx: context.NewRootContext(width, height, &context.Margins{
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
		}),
	}
}

func (d *document) ForceAddPage(pages ...Page) {
	d.pages = append(d.pages, pages...)
}

func (d *document) Add(rows ...Row) {
	d.rows = append(d.rows, rows...)
}

func (d *document) Generate() error {
	d.ctx.Print(d._type)
	ctx := d.ctx.WithDimension(d.ctx.MaxWidth(), d.ctx.MaxHeight())

	for _, row := range d.rows {
		row.Render(d.fpdf, ctx)
	}

	for _, page := range d.pages {
		page.Render(d.fpdf, ctx)
	}

	return d.fpdf.OutputFileAndClose(d.file)
}

func (d *document) GetStructure() *tree.Node[Structure] {
	str := Structure{
		Type:  string(d._type),
		Value: d.file,
	}
	node := tree.NewNode(0, str)

	for _, r := range d.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}
