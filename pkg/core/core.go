package core

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/config"
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
	Render(provider Provider, cell Cell)
}

type Col interface {
	Node
	Add(components ...Component) Col
	GetSize() int
	WithStyle(style *props.Cell) Col
	Render(provider Provider, cell Cell, createCell bool)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
	WithStyle(style *props.Cell) Row
	Render(provider Provider, cell Cell)
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int, total int)
	Render(provider Provider, cell Cell)
}
