package v2

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/jung-kurt/gofpdf"
)

type Maroto interface {
	Add(component ...Component)
	Generate(file string) error
}

type document struct {
	ctx        context.Context
	_type      DocumentType
	fpdf       fpdf.Fpdf
	components []Component
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
		_type: Document,
		ctx: context.NewRootContext(width, height, &context.Margins{
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
		}),
	}
}

func (d *document) Add(components ...Component) {
	for _, component := range components {
		if d._type.Accept(component.GetType()) {
			d.components = append(d.components, component)
		}
	}
}

func (d *document) Generate(file string) error {
	d.ctx.Print(d._type)
	ctx := d.ctx.WithDimension(d.ctx.MaxWidth(), d.ctx.MaxHeight())

	for _, component := range d.components {
		component.Render(d.fpdf, ctx)
	}

	return d.fpdf.OutputFileAndClose(file)
}
