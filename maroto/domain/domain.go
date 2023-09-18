package domain

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/v2/internal"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/metrics"
	"github.com/johnfercher/v2/maroto/props"
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
	SetConfig(config *config.Maroto)
	GetStructure() *tree.Node[Structure]
}

type Component interface {
	Node
	Render(provider Provider, cell internal.Cell)
}

type Col interface {
	Node
	Add(components ...Component) Col
	GetSize() int
	WithStyle(style *props.Style) Col
	Render(provider Provider, cell internal.Cell, createCell bool)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
	WithStyle(style *props.Style) Row
	Render(provider Provider, cell internal.Cell)
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int)
	Render(provider Provider, cell internal.Cell)
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
