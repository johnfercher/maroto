package gofpdf

import (
	"bytes"

	"github.com/johnfercher/maroto/v2/internal/math"

	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/providers"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf/cellwriter"

	"github.com/jung-kurt/gofpdf"
)

var defaultErrorColor = &props.Font{
	Family: fontfamily.Arial,
	Style:  fontstyle.Bold,
	Size:   10,
	Color: &props.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	},
}

type gofpdfProvider struct {
	fpdf       *gofpdf.Fpdf
	math       math.Math
	font       internal.Font
	text       internal.Text
	signature  internal.Signature
	code       internal.Code
	image      internal.Image
	line       internal.Line
	imageCache cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *config.Config
}

func New(cfg *config.Config, options ...providers.ProviderOption) core.Provider {
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
		fpdf.AddUTF8FontFromBytes(font.Family, string(font.Style), font.Bytes)
	}

	fpdf.SetMargins(cfg.Margins.Left, cfg.Margins.Top, cfg.Margins.Right)
	fpdf.AddPage()

	font := internal.NewFont(fpdf, cfg.DefaultFont.Size, cfg.DefaultFont.Family, cfg.DefaultFont.Style)
	math := math.NewMath()
	text := internal.NewText(fpdf, math, font)
	signature := internal.NewSignature(fpdf, math, text)
	code := internal.NewCode(fpdf, math)
	image := internal.NewImage(fpdf, math)
	line := internal.NewLine(fpdf)
	cellWriter := cellwriter.NewBuilder().Build(fpdf)

	provider := &gofpdfProvider{
		fpdf:       fpdf,
		math:       math,
		font:       font,
		text:       text,
		signature:  signature,
		code:       code,
		image:      image,
		line:       line,
		cellWriter: cellWriter,
		cfg:        cfg,
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

func (g *gofpdfProvider) AddText(text string, cell *core.Cell, prop *props.Text) {
	g.text.Add(text, cell, prop)
}

func (g *gofpdfProvider) AddLine(cell *core.Cell, prop *props.Line) {
	g.line.Add(cell, prop)
}

func (g *gofpdfProvider) AddSignature(text string, cell *core.Cell, prop *props.Text) {
	g.signature.AddSpaceFor(text, cell, prop)
}

func (g *gofpdfProvider) AddMatrixCode(code string, cell *core.Cell, prop *props.Rect) {
	g.code.AddDataMatrix(code, cell, g.cfg.Margins, prop)
}

func (g *gofpdfProvider) AddQrCode(code string, cell *core.Cell, rect *props.Rect) {
	g.code.AddQr(code, cell, g.cfg.Margins, rect)
}

func (g *gofpdfProvider) AddBarCode(code string, cell *core.Cell, prop *props.Barcode) {
	err := g.code.AddBar(code, cell, g.cfg.Margins, prop)
	if err != nil {
		textProp := &props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to render code", cell, textProp)
	}
}

func (g *gofpdfProvider) AddImage(file string, cell *core.Cell, prop *props.Rect, extension extension.Type) {
	img, err := g.imageCache.Load(file, extension)
	if err != nil {
		textProp := &props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to load image from file", cell, textProp)
		return
	}

	err = g.image.AddFromBase64(img.Value, cell, g.cfg.Margins, prop, img.Extension)
	if err != nil {
		textProp := &props.Text{}
		textProp.MakeValid(defaultErrorColor)
		g.fpdf.ClearError()
		g.AddText("Failed to load image from file", cell, textProp)
	}
}

func (g *gofpdfProvider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *gofpdfProvider) SetProtection(protection *config.Protection) {
	if protection == nil {
		return
	}

	g.fpdf.SetProtection(byte(protection.Type), protection.UserPassword, protection.OwnerPassword)
}

func (g *gofpdfProvider) SetMetadata(metadata *config.Metadata) {
	if metadata == nil {
		return
	}

	if metadata.Author != nil {
		g.fpdf.SetAuthor(metadata.Author.Text, metadata.Author.UTF8)
	}

	if metadata.Creator != nil {
		g.fpdf.SetCreator(metadata.Creator.Text, metadata.Creator.UTF8)
	}

	if metadata.Subject != nil {
		g.fpdf.SetSubject(metadata.Subject.Text, metadata.Subject.UTF8)
	}

	if metadata.Title != nil {
		g.fpdf.SetTitle(metadata.Title.Text, metadata.Title.UTF8)
	}

	if !metadata.CreationDate.IsZero() {
		g.fpdf.SetCreationDate(metadata.CreationDate)
	}
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

func (g *gofpdfProvider) CreateCol(width, height float64, config *config.Config, prop *props.Cell) {
	g.cellWriter.Apply(width, height, config, prop)
}

func (g *gofpdfProvider) SetCompression(compression bool) {
	g.fpdf.SetCompression(compression)
}
