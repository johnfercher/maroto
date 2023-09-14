package providers

import (
	"bytes"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/size"
	"github.com/jung-kurt/gofpdf"
)

type gofpdfProvider struct {
	fpdf      *gofpdf.Fpdf
	math      internal.Math
	font      internal.Font
	text      internal.Text
	signature internal.Signature
	code      internal.Code
	image     internal.Image
}

func NewGofpdf(pageSize size.PageSize) domain.Provider {
	width, height := size.GetDimensions(pageSize)

	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		FontDirStr:     "",
		Size: gofpdf.SizeType{
			Wd: width,
			Ht: height,
		},
	})

	fpdf.SetFont("Arial", "B", 16)
	fpdf.AddPage()

	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)
	signature := internal.NewSignature(fpdf, math, text)
	code := internal.NewCode(fpdf, math)
	image := internal.NewImage(fpdf, math)

	return &gofpdfProvider{
		fpdf:      fpdf,
		math:      math,
		font:      font,
		text:      text,
		signature: signature,
		code:      code,
		image:     image,
	}
}

func (g *gofpdfProvider) GetDimensions() (width float64, height float64) {
	return g.fpdf.GetPageSize()
}

func (g *gofpdfProvider) GetMargins() (left float64, top float64, right float64, bottom float64) {
	return g.fpdf.GetMargins()
}

func (g *gofpdfProvider) AddText(text string, cell internal.Cell, prop props.Text) {
	g.text.Add(text, cell, prop)
}

func (g *gofpdfProvider) AddSignature(text string, cell internal.Cell, prop props.Text) {
	g.signature.AddSpaceFor(text, cell, prop)
}

func (g *gofpdfProvider) AddMatrixCode(code string, cell internal.Cell, prop props.Rect) {
	g.code.AddDataMatrix(code, cell, prop)
}

func (g *gofpdfProvider) AddQrCode(code string, cell internal.Cell, rect props.Rect) {
	g.code.AddQr(code, cell, rect)
}

func (g *gofpdfProvider) AddBarCode(code string, cell internal.Cell, prop props.Barcode) {
	err := g.code.AddBar(code, cell, prop)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(consts.Arial)
		g.fpdf.ClearError()
		g.AddText("Failed to render code", cell, textProp)
	}
}

func (g *gofpdfProvider) AddImageFromBase64(base64 string, cell internal.Cell, prop props.Rect, extension consts.Extension) {
	err := g.image.AddFromBase64(base64, cell, prop, extension)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(consts.Arial)
		g.fpdf.ClearError()
		g.AddText("Failed to render image from base64", cell, textProp)
	}
}

func (g *gofpdfProvider) AddImageFromFile(file string, cell internal.Cell, prop props.Rect) {
	err := g.image.AddFromFile(file, cell, prop)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(consts.Arial)
		g.fpdf.ClearError()
		g.AddText("Failed to render image from file", cell, textProp)
	}
}

func (g *gofpdfProvider) CreateCol(width, height float64) {
	g.fpdf.CellFormat(width, height, "", "1", 0, "C", false, 0, "")
}

func (g *gofpdfProvider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *gofpdfProvider) Generate(file string) error {
	return g.fpdf.OutputFileAndClose(file)
}

func (g *gofpdfProvider) GenerateAndOutput() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	err := g.fpdf.Output(&buffer)
	return buffer, err
}
