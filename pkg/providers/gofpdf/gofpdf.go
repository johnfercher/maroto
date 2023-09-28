package gofpdf

import (
	"bytes"
	"encoding/base64"
	"strings"

	"github.com/johnfercher/maroto/v2/internal/code"

	"github.com/johnfercher/maroto/v2/pkg/merror"

	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf/cellwriter"

	"github.com/jung-kurt/gofpdf"
)

type gofpdfProvider struct {
	fpdf       *gofpdf.Fpdf
	font       core.Font
	text       core.Text
	signature  core.Signature
	code       core.Code
	image      core.Image
	line       core.Line
	cache      cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *config.Config
}

func New(cfg *config.Config, cache cache.Cache) core.Provider {
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
	math := math.New()
	code := code.NewCode()
	text := internal.NewText(fpdf, math, font)
	signature := internal.NewSignature(fpdf, math, text)
	image := internal.NewImage(fpdf, math)
	line := internal.NewLine(fpdf)
	cellWriter := cellwriter.NewBuilder().Build(fpdf)

	provider := &gofpdfProvider{
		fpdf:       fpdf,
		font:       font,
		text:       text,
		signature:  signature,
		code:       code,
		image:      image,
		line:       line,
		cellWriter: cellWriter,
		cfg:        cfg,
		cache:      cache,
	}

	return provider
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
	bytes, err := g.cache.GetCode(code, "matrixcode")
	if err != nil {
		bytes, err = g.code.GenDataMatrix(code)
	}

	if err != nil {
		g.text.Add("could not generate matrixcode", cell, merror.DefaultErrorText)
		return
	}

	g.cache.SaveCode(code, "matrixcode", bytes)
	err = g.image.Add(bytes, cell, g.cfg.Margins, prop, extension.Jpg)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add matrixcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *gofpdfProvider) AddQrCode(code string, cell *core.Cell, prop *props.Rect) {
	bytes, err := g.cache.GetCode(code, "qrcode")
	if err != nil {
		bytes, err = g.code.GenQr(code)
	}

	if err != nil {
		g.text.Add("could not generate qrcode", cell, merror.DefaultErrorText)
		return
	}

	g.cache.SaveCode(code, "qrcode", bytes)
	err = g.image.Add(bytes, cell, g.cfg.Margins, prop, extension.Jpg)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add qrcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *gofpdfProvider) AddBarCode(code string, cell *core.Cell, prop *props.Barcode) {
	bytes, err := g.cache.GetCode(code, "barcode")
	if err != nil {
		bytes, err = g.code.GenBar(code, cell, prop)
	}

	if err != nil {
		g.text.Add("could not generate barcode", cell, merror.DefaultErrorText)
		return
	}

	g.cache.SaveCode(code, "barcode", bytes)
	err = g.image.Add(bytes, cell, g.cfg.Margins, prop.ToRectProp(), extension.Jpg)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add barcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *gofpdfProvider) AddImageFromFile(file string, cell *core.Cell, prop *props.Rect) {
	extensionStr := strings.Split(file, ".")[1]
	image, err := g.cache.GetImage(file, extensionStr)
	if err != nil {
		err = g.cache.LoadImage(file, extensionStr)
	} else {
		g.AddImageFromBytes(image.Bytes, cell, prop, extension.Type(extensionStr))
		return
	}

	if err != nil {
		g.text.Add("could not load image", cell, merror.DefaultErrorText)
		return
	}

	image, err = g.cache.GetImage(file, extensionStr)
	if err != nil {
		g.text.Add("could not load image", cell, merror.DefaultErrorText)
		return
	}

	g.AddImageFromBytes(image.Bytes, cell, prop, extension.Type(extensionStr))
}

func (g *gofpdfProvider) AddImageFromBase64(base64string string, cell *core.Cell, prop *props.Rect, extension extension.Type) {
	bytes, err := base64.StdEncoding.DecodeString(base64string)
	if err != nil {
		g.text.Add("could not parse image from base64", cell, merror.DefaultErrorText)
		return
	}

	g.AddImageFromBytes(bytes, cell, prop, extension)
}

func (g *gofpdfProvider) AddImageFromBytes(bytes []byte, cell *core.Cell, prop *props.Rect, extension extension.Type) {
	err := g.image.Add(bytes, cell, g.cfg.Margins, prop, extension)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add image to document", cell, merror.DefaultErrorText)
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

func (g *gofpdfProvider) GenerateBytes() ([]byte, error) {
	var buffer bytes.Buffer
	err := g.fpdf.Output(&buffer)

	return buffer.Bytes(), err
}

func (g *gofpdfProvider) SetCache(cache cache.Cache) {
	g.cache = cache
}

func (g *gofpdfProvider) CreateCol(width, height float64, config *config.Config, prop *props.Cell) {
	g.cellWriter.Apply(width, height, config, prop)
}

func (g *gofpdfProvider) SetCompression(compression bool) {
	g.fpdf.SetCompression(compression)
}
