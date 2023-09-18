package domain

import (
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64, config *config.Maroto, style *props.Style)

	// Features
	AddText(text string, cell internal.Cell, prop props.Text)
	AddSignature(text string, cell internal.Cell, prop props.Text)
	AddMatrixCode(code string, cell internal.Cell, prop props.Rect)
	AddQrCode(code string, cell internal.Cell, rect props.Rect)
	AddBarCode(code string, cell internal.Cell, prop props.Barcode)
	AddImage(value string, cell internal.Cell, prop props.Rect, extension consts.Extension)

	// General
	GetDimensions() (width float64, height float64)
	GetMargins() (left float64, top float64, right float64, bottom float64)
	GenerateFile(file string) error
	GenerateBytes() ([]byte, error)
	SetCache(cache cache.Cache)
}
