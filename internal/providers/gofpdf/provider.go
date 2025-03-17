package gofpdf

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"

	"github.com/johnfercher/maroto/v2/internal/cache"
	"github.com/johnfercher/maroto/v2/internal/merror"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type provider struct {
	fpdf       gofpdfwrapper.Fpdf
	font       core.Font
	text       core.Text
	code       core.Code
	image      core.Image
	line       core.Line
	heatMap    core.HeatMap
	timeSeries core.TimeSeries
	cache      cache.Cache
	cellWriter cellwriter.CellWriter
	cfg        *entity.Config
}

// New is the constructor of provider for gofpdf
func New(dep *Dependencies) core.Provider {
	return &provider{
		fpdf:       dep.Fpdf,
		font:       dep.Font,
		text:       dep.Text,
		code:       dep.Code,
		image:      dep.Image,
		line:       dep.Line,
		heatMap:    dep.HeatMap,
		timeSeries: dep.TimeSeries,
		cellWriter: dep.CellWriter,
		cfg:        dep.Cfg,
		cache:      dep.Cache,
	}
}

func (g *provider) AddText(text string, cell *entity.Cell, prop *props.Text) {
	g.text.Add(text, cell, prop)
}

func (g *provider) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	return g.text.GetLinesQuantity(text, textProp, colWidth)
}

func (g *provider) GetFontHeight(prop *props.Font) float64 {
	return g.font.GetHeight(prop.Family, prop.Style, prop.Size)
}

func (g *provider) AddLine(cell *entity.Cell, prop *props.Line) {
	g.line.Add(cell, prop)
}

func (g *provider) AddMatrixCode(code string, cell *entity.Cell, prop *props.Rect) {
	img, err := g.loadCode(code, "matrix-code-", g.code.GenDataMatrix)
	if err != nil {
		g.text.Add("could not generate matrixcode", cell, merror.DefaultErrorText)
		return
	}

	err = g.image.Add(img, cell, g.cfg.Margins, prop, extension.Png, false)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add matrixcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddQrCode(code string, cell *entity.Cell, prop *props.Rect) {
	img, err := g.loadCode(code, "qr-code-", g.code.GenQr)
	if err != nil {
		g.text.Add("could not generate qrcode", cell, merror.DefaultErrorText)
		return
	}

	err = g.image.Add(img, cell, g.cfg.Margins, prop, extension.Png, false)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add qrcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddBarCode(code string, cell *entity.Cell, prop *props.Barcode) {
	image, err := g.cache.GetImage(g.getBarcodeImageName(fmt.Sprintf("bar-code-%s", code), prop), extension.Png)
	if err != nil {
		image, err = g.code.GenBar(code, cell, prop)
	}
	if err != nil {
		g.text.Add("could not generate barcode", cell, merror.DefaultErrorText)
		return
	}

	g.cache.AddImage(g.getBarcodeImageName(fmt.Sprintf("bar-code-%s", code), prop), image)
	err = g.image.Add(image, cell, g.cfg.Margins, prop.ToRectProp(), extension.Png, false)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add barcode to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddImageFromFile(file string, cell *entity.Cell, prop *props.Rect) {
	extensionStr := strings.ToLower(strings.TrimPrefix(filepath.Ext(file), "."))
	image, err := g.loadImage(file, extensionStr)
	if err != nil {
		g.text.Add("could not load image", cell, merror.DefaultErrorText)
		return
	}

	g.AddImageFromBytes(image.Bytes, cell, prop, extension.Type(extensionStr))
}

func (g *provider) AddImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type) {
	img, err := FromBytes(bytes, extension)
	if err != nil {
		g.text.Add("could not parse image bytes", cell, merror.DefaultErrorText)
		return
	}

	err = g.image.Add(img, cell, g.cfg.Margins, prop, extension, false)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add image to document", cell, merror.DefaultErrorText)
	}
}

func (g *provider) AddBackgroundImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type) {
	img, err := FromBytes(bytes, extension)
	if err != nil {
		g.text.Add("could not parse image bytes", cell, merror.DefaultErrorText)
		return
	}

	err = g.image.Add(img, cell, g.cfg.Margins, prop, extension, true)
	if err != nil {
		g.fpdf.ClearError()
		g.text.Add("could not add image to document", cell, merror.DefaultErrorText)
	}
	g.fpdf.SetHomeXY()
}

func (g *provider) CreateRow(height float64) {
	g.fpdf.Ln(height)
}

func (g *provider) SetProtection(protection *entity.Protection) {
	if protection == nil {
		return
	}

	g.fpdf.SetProtection(byte(protection.Type), protection.UserPassword, protection.OwnerPassword)
}

func (g *provider) SetMetadata(metadata *entity.Metadata) {
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

	if metadata.CreationDate != nil {
		g.fpdf.SetCreationDate(*metadata.CreationDate)
	}

	if metadata.KeywordsStr != nil {
		g.fpdf.SetKeywords(metadata.KeywordsStr.Text, metadata.KeywordsStr.UTF8)
	}
}

func (g *provider) AddTimeSeries(timeSeriesList []entity.TimeSeries, cell *entity.Cell) {
	g.timeSeries.Add(timeSeriesList, cell, g.cfg.Margins)
}

// GetDimensionsByImage is responsible for obtaining the dimensions of an image
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByImage(file string) (*entity.Dimensions, error) {
	extensionStr := strings.ToLower(strings.TrimPrefix(filepath.Ext(file), "."))
	img, err := g.loadImage(file, extensionStr)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Type(extensionStr))

	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}
	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

// GetDimensionsByImageByte is responsible for obtaining the dimensions of an image
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByImageByte(bytes []byte, extension extension.Type) (*entity.Dimensions, error) {
	img, err := FromBytes(bytes, extension)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension)
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}
	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

// GetDimensionsByMatrixCode is responsible for obtaining the dimensions of an MatrixCode
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByMatrixCode(code string) (*entity.Dimensions, error) {
	img, err := g.loadCode(code, "matrix-code-", g.code.GenDataMatrix)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Png)

	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}
	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

// GetDimensionsByQrCode is responsible for obtaining the dimensions of an QrCode
// If the image cannot be loaded, an error is returned
func (g *provider) GetDimensionsByQrCode(code string) (*entity.Dimensions, error) {
	img, err := g.loadCode(code, "qr-code-", g.code.GenQr)
	if err != nil {
		return nil, err
	}

	imgInfo, _ := g.image.GetImageInfo(img, extension.Png)
	if imgInfo == nil {
		return nil, errors.New("could not read image options, maybe path/name is wrong")
	}
	return &entity.Dimensions{Width: imgInfo.Width(), Height: imgInfo.Height()}, nil
}

func (g *provider) GenerateBytes() ([]byte, error) {
	var buffer bytes.Buffer
	err := g.fpdf.Output(&buffer)

	return buffer.Bytes(), err
}

func (g *provider) CreateCol(width, height float64, config *entity.Config, prop *props.Cell) {
	g.cellWriter.Apply(width, height, config, prop)
}

func (g *provider) SetCompression(compression bool) {
	g.fpdf.SetCompression(compression)
}

func (g *provider) AddHeatMap(heatMap [][]int, cell *entity.Cell, prop *props.HeatMap) {
	g.heatMap.Add(heatMap, cell, g.cfg.Margins, prop)
}

func (g *provider) getBarcodeImageName(code string, prop *props.Barcode) string {
	if prop == nil {
		return code + string(barcode.Code128)
	}

	return code + string(prop.Type)
}

// loadImage is responsible for loading an codes
func (g *provider) loadCode(code, codeType string, generate func(code string) (*entity.Image, error)) (*entity.Image, error) {
	image, err := g.cache.GetImage(codeType+code, extension.Png)
	if err != nil {
		image, err = generate(code)
	} else {
		return image, nil
	}
	if err != nil {
		return nil, err
	}
	g.cache.AddImage(codeType+code, image)

	return image, nil
}

// loadImage is responsible for loading an image
func (g *provider) loadImage(file, extensionStr string) (*entity.Image, error) {
	image, err := g.cache.GetImage(file, extension.Type(extensionStr))

	if err == nil {
		return image, err
	}

	if err = g.cache.LoadImage(file, extension.Type(extensionStr)); err != nil {
		return nil, err
	}

	return g.cache.GetImage(file, extension.Type(extensionStr))
}
