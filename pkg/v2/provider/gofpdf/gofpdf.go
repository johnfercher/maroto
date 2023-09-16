package gofpdf

import (
	"bytes"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/cache"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/providers"
	"github.com/jung-kurt/gofpdf"
)

type gofpdfProvider struct {
	fpdf       *gofpdf.Fpdf
	math       internal.Math
	font       internal.Font
	text       internal.Text
	signature  internal.Signature
	code       internal.Code
	image      internal.Image
	imageCache cache.Cache
}

func New(maroto *config.Maroto, options ...providers.ProviderOption) domain.Provider {
	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		FontDirStr:     "",
		Size: gofpdf.SizeType{
			Wd: maroto.Dimensions.Width,
			Ht: maroto.Dimensions.Height,
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

	provider := &gofpdfProvider{
		fpdf:      fpdf,
		math:      math,
		font:      font,
		text:      text,
		signature: signature,
		code:      code,
		image:     image,
	}

	for _, option := range options {
		option(provider)
	}

	return provider
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

func (g *gofpdfProvider) AddImage(file string, cell internal.Cell, prop props.Rect, extension consts.Extension) {
	img, err := g.imageCache.Load(file, extension)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(consts.Arial)
		g.fpdf.ClearError()
		g.AddText("Failed to load image from file", cell, textProp)
	}

	g.image.AddFromBase64(img.Value, cell, prop, img.Extension)
}

func (g *gofpdfProvider) CreateCol(width, height float64, config *config.Maroto) {
	border := "0"
	if config.Debug {
		border = "1"
	}

	g.fpdf.CellFormat(width, height, "", border, 0, "C", false, 0, "")
}

func (g *gofpdfProvider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *gofpdfProvider) GenerateFile(file string) error {
	return g.fpdf.OutputFileAndClose(file)
}

func (g *gofpdfProvider) GenerateBytes() ([]byte, error) {
	var buffer bytes.Buffer
	err := g.fpdf.Output(&buffer)

	return buffer.Bytes(), err
}

func (g *gofpdfProvider) SetCache(cache cache.Cache) {
	g.imageCache = cache
}
