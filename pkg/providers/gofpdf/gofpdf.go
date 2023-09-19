package gofpdf

import (
	"bytes"

	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/color"
	"github.com/johnfercher/maroto/v2/pkg/core/context"

	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/providers"

	"github.com/jung-kurt/gofpdf"
)

var defaultErrorColor = &props.Font{
	Family: fontfamily.Arial,
	Style:  fontstyle.Bold,
	Size:   10,
	Color: &color.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	},
}

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

func New(cfg *config.Maroto, options ...providers.ProviderOption) core.Provider {
	fpdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		FontDirStr:     "",
		Size: gofpdf.SizeType{
			Wd: cfg.Dimensions.Width,
			Ht: cfg.Dimensions.Height,
		},
	})

	for _, font := range cfg.CustomFonts {
		fpdf.AddUTF8Font(font.Family, string(font.Style), font.File)
	}

	fpdf.SetMargins(cfg.Margins.Left, cfg.Margins.Top, cfg.Margins.Right)
	fpdf.AddPage()

	font := internal.NewFont(fpdf, cfg.Font.Size, cfg.Font.Family, cfg.Font.Style)
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

func (g *gofpdfProvider) AddText(text string, cell context.Cell, prop props.Text) {
	g.text.Add(text, cell, prop)
}

func (g *gofpdfProvider) AddSignature(text string, cell context.Cell, prop props.Text) {
	g.signature.AddSpaceFor(text, cell, prop)
}

func (g *gofpdfProvider) AddMatrixCode(code string, cell context.Cell, prop props.Rect) {
	g.code.AddDataMatrix(code, cell, prop)
}

func (g *gofpdfProvider) AddQrCode(code string, cell context.Cell, rect props.Rect) {
	g.code.AddQr(code, cell, rect)
}

func (g *gofpdfProvider) AddBarCode(code string, cell context.Cell, prop props.Barcode) {
	err := g.code.AddBar(code, cell, prop)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to render code", cell, textProp)
	}
}

func (g *gofpdfProvider) AddImage(file string, cell context.Cell, prop props.Rect, extension extension.Type) {
	img, err := g.imageCache.Load(file, extension)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to load image from file", cell, textProp)
		return
	}

	err = g.image.AddFromBase64(img.Value, cell, prop, img.Extension)
	if err != nil {
		textProp := props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to load image from file", cell, textProp)
	}
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

func (g *gofpdfProvider) CreateCol(width, height float64, config *config.Maroto, style *props.Style) {
	if style == nil {
		g.createStandard(width, height, config)
		return
	}

	g.createCustom(width, height, config, style)
}

func (g *gofpdfProvider) createStandard(width, height float64, config *config.Maroto) {
	border := "0"
	if config.Debug {
		border = "1"
	}

	g.fpdf.CellFormat(width, height, "", border, 0, "C", false, 0, "")
}

func (g *gofpdfProvider) createCustom(width, height float64, config *config.Maroto, style *props.Style) {
	bd := style.Border
	if config.Debug {
		bd = border.Full
	}

	fill := false
	if style.BackgroundColor != nil {
		g.fpdf.SetFillColor(style.BackgroundColor.Red, style.BackgroundColor.Green, style.BackgroundColor.Blue)
		fill = true
	}

	if fill && style.BorderColor != nil {
		g.fpdf.SetDrawColor(style.BorderColor.Red, style.BorderColor.Green, style.BorderColor.Blue)
	}

	g.fpdf.CellFormat(width, height, "", string(bd), 0, "C", fill, 0, "")

	if fill {
		white := color.NewWhite()
		g.fpdf.SetFillColor(white.Red, white.Green, white.Blue)
	}

	if fill && style.BorderColor != nil {
		black := color.NewBlack()
		g.fpdf.SetDrawColor(black.Red, black.Green, black.Blue)
	}
}
