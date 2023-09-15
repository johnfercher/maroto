package domain

import (
	"bytes"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

type ProviderType string

const (
	Gofpdf ProviderType = "gofpdf"
	HTML   ProviderType = "html"
)

type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64)

	// Features
	AddText(text string, cell internal.Cell, prop props.Text)
	AddSignature(text string, cell internal.Cell, prop props.Text)
	AddMatrixCode(code string, cell internal.Cell, prop props.Rect)
	AddQrCode(code string, cell internal.Cell, rect props.Rect)
	AddBarCode(code string, cell internal.Cell, prop props.Barcode)
	AddImageFromBase64(base64 string, cell internal.Cell, prop props.Rect, extension consts.Extension)
	AddImageFromFile(file string, cell internal.Cell, prop props.Rect)

	// General
	GetDimensions() (width float64, height float64)
	GetMargins() (left float64, top float64, right float64, bottom float64)
	Generate(file string) error
	GenerateAndOutput() (bytes.Buffer, error)
}
