package providers

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/jung-kurt/gofpdf"
)

type gofpdfProvider struct {
	fpdf *gofpdf.Fpdf
	math internal.Math
	font internal.Font
	text internal.Text
}

func NewGofpdf() domain.Provider {
	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        "A4",
		FontDirStr:     "",
	})

	fpdf.SetFont("Arial", "B", 16)
	fpdf.AddPage()

	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)

	return &gofpdfProvider{
		fpdf: fpdf,
		math: math,
		font: font,
		text: text,
	}
}

func (g *gofpdfProvider) GetDimensions() (width float64, height float64) {
	return g.fpdf.GetPageSize()
}

func (g *gofpdfProvider) GetMargins() (left float64, top float64, right float64, bottom float64) {
	return g.fpdf.GetMargins()
}

func (g *gofpdfProvider) AddText(text string, cell internal.Cell, prop props.Text) {
	g.text.Add(
		text,
		cell,
		prop)
}

func (g *gofpdfProvider) CreateCol(width, height float64) {
	g.fpdf.CellFormat(width, height, "", "1", 0, "C", false, 0, "")
}

func (g *gofpdfProvider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *gofpdfProvider) Generate(file string) error {
	return g.fpdf.OutputFileAndClose(file)
	return g.fpdf.OutputFileAndClose(file)
}
