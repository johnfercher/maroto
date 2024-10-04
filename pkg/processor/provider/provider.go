package provider

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
)

type Provider interface {
	GeneratePdf() ([]byte, error)
	ConfigureBuilder(builder builder.Builder) Provider
	RegisterHeader(rows ...core.Row) Provider
	RegisterFooter(rows ...core.Row) Provider
	CreatePage(components ...core.Row) core.Page
	CreateRow(components ...core.Col) core.Row
	CreateCol(size int, components ...core.Component) core.Col
	CreateText(value string, props props.TextProps) core.Component
	CreateBarCode(value string, props props.BarCodeProps) core.Component
}
