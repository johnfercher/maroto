package core

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/context"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Maroto interface {
	RegisterHeader(rows ...Row) error
	RegisterFooter(rows ...Row) error
	AddRows(rows ...Row)
	AddRow(rowHeight float64, cols ...Col) Row
	ForceAddPage(pages ...Page)
	GetStructure() *tree.Node[Structure]
	Generate() (Document, error)
}

type Document interface {
	GetBytes() []byte
	GetBase64() string
	Save(file string) error
	GetReport() *metrics.Report
}

type Node interface {
	SetConfig(config *config.Config)
	GetStructure() *tree.Node[Structure]
}

type Component interface {
	Node
	Render(provider Provider, cell context.Cell)
}

type Col interface {
	Node
	Add(components ...Component) Col
	GetSize() int
	WithStyle(style *props.Cell) Col
	Render(provider Provider, cell context.Cell, createCell bool)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
	WithStyle(style *props.Cell) Row
	Render(provider Provider, cell context.Cell)
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int)
	Render(provider Provider, cell context.Cell)
}

type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64, config *config.Config, style *props.Cell)

	// Features
	AddText(text string, cell context.Cell, prop props.Text)
	AddSignature(text string, cell context.Cell, prop props.Text)
	AddMatrixCode(code string, cell context.Cell, prop props.Rect)
	AddQrCode(code string, cell context.Cell, rect props.Rect)
	AddBarCode(code string, cell context.Cell, prop props.Barcode)
	AddImage(value string, cell context.Cell, prop props.Rect, extension extension.Type)

	// General
	GetDimensions() (width float64, height float64)
	GetMargins() (left float64, top float64, right float64, bottom float64)
	GenerateFile(file string) error
	GenerateBytes() ([]byte, error)
	SetCache(cache cache.Cache)
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
