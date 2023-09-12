package v2

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/jung-kurt/gofpdf"
)

type Maroto interface {
	Add(component ...Component)
	Generate(file string) error
}

type document struct {
	ctx        Context
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

	width, height := fpdf.GetPageSize()
	left, top, right, bottom := fpdf.GetMargins()
	fpdf.AddPage()

	return &document{
		fpdf:  fpdf,
		_type: Document,
		ctx: NewRootContext(width, height, &Margins{
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
	for _, component := range d.components {
		component.Render(d.fpdf, &d.ctx)
	}

	return d.fpdf.OutputFileAndClose(file)
}
